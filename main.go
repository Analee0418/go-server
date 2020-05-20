package main

import (
	"bytes"
	"compress/flate"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"runtime/debug"
	"syscall"
	"time"

	"com.lueey.shop/common"
	"com.lueey.shop/config"
	"com.lueey.shop/handler"
	"com.lueey.shop/model"
	avro "com.lueey.shop/protocol"
	"com.lueey.shop/utils"
	"github.com/actgardner/gogen-avro/compiler"
	"github.com/actgardner/gogen-avro/vm"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
}

func main() {
	common.ServerCategory = common.SERVER_CATEGORY_HALL

	defer func() {
		if x := recover(); x != nil {
			debug.PrintStack()
			log.Println("FATAL: caught panic in main()", x)
		}
	}()

	runtime.GOMAXPROCS(runtime.NumCPU())

	config.InitDBConfig()
	utils.InitRedisDB()

	config.Init()

	model.InitAuctionGoods()
	model.InitCustomer()
	model.InitRoom()
	model.TCPInitGlobal()

	model.PostInitAuctionGoods()
	model.PostInitCustoemr()
	model.PostInitRoom()

	model.TCPServerInit()

	// crontab 任务
	startTimer(func(now int64) {
		// session
		model.ReleaseSessionCache(now)

		// server refresh
		model.TCPServerInstance.TCPServerRefresh(now)
	})

	// 接收世界状态
	go model.TCPGlobalReceiveGlobalState()

	// 接收世界广播
	go model.OnBroadcastToGlobal()

	// 退出
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2) // 监听信号
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
				log.Println("\033[1;33m[WARNING] \033[0mServer program exit:", s)
				model.TCPServerInstance.TCPServerUpdateStatus(model.ServerDeactivate, false)
				utils.HDelRedis("hallserver##startup", model.TCPServerInstance.ID)
				log.Println("[INFO] Notify HTTPServer")
				os.Exit(0)
			case syscall.SIGUSR1:
				log.Println("[INFO] usr1", s)
			case syscall.SIGUSR2:
				log.Println("[INFO] usr2", s)
			default:
				log.Println("[INFO] Other", s)
			}
		}
	}()

	startTCPServer()

}

const (
	buffSize int    = 1024 * 128
	pORT     string = "1234"
)

func handleConnection(conn net.Conn) {
	defer func() {
		if x := recover(); x != nil {
			debug.PrintStack()
			log.Println("\033[1;31m[ERROR] \033[0mcaught panic in handleConnection", x)
		}
		time.Sleep(10 * time.Microsecond)
	}()

	ip := conn.RemoteAddr().String()

	_ret := make([]byte, 4)
	n, e := conn.Read(_ret)
	if e != nil {
		if e == io.EOF { // 客户端主动断开
			log.Println("\033[1;33m[WARNING] \033[0m!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! discover client disconnect.", ip)
			model.OnClientDisconnect(conn)
			return
		}
	}
	dataLen := int(binary.LittleEndian.Uint32(_ret))
	if config.DEBUG {
		log.Println("[INFO] ->->->->->->->->->->->-> length", dataLen, n, ip, e)
	}
	if dataLen >= 2097152 {
		log.Println("\033[1;31m[ERROR] \033[0mToo large data packets! ->->->->->->->->->->->-> ", ip)
		conn.Close()
		return
	}

	ulimitBuffer := bytes.NewBuffer(nil)
	remain := dataLen
	for {
		l := buffSize
		if l < dataLen {
			l = dataLen
		}
		if l > remain {
			l = remain
		}
		pkg := make([]byte, l)
		n, _ := conn.Read(pkg)
		if config.DEBUG {
			log.Printf("[DEBUG] len: %d, read len: %d, remain: %d\n", l, n, remain)
		}
		if n == 0 {
			break
		}
		ulimitBuffer.Write(pkg[:n])

		remain = dataLen - len(ulimitBuffer.Bytes())
		if remain <= 0 {
			break
		}
	}

	if config.DEBUG {
		log.Printf("[DEBUG] ->->->->->->->->->->->-> read all length: %d", ulimitBuffer.Len())
	}

	message := avro.Message{}
	deser, err := compiler.CompileSchemaBytes([]byte(message.Schema()), []byte(message.Schema()))
	if err != nil {
		log.Printf("\033[1;31m[ERROR] \033[0mcannot found schema compiler. IP: %s, Err: %v\n\n", ip, err)
		return
	}
	buffer := bytes.NewBuffer(ulimitBuffer.Bytes())
	var reader = flate.NewReader(buffer)
	errs := vm.Eval(reader, deser, &message)
	if errs != nil {
		log.Printf("\033[1;31m[ERROR] \033[0mcannot Decompress/ Decode data. IP: %s, Err: %v\n\n", ip, errs)
		return
	}

	var body string = ""
	lang, err := json.MarshalIndent(message, "", "   ")
	if err == nil {
		body = string(lang)
	}
	log.Printf("[INFO] ACTION=%s, DATALEN=%d, REMOTE_IP=%s, BODY=%s\n", message.Action, ulimitBuffer.Len(), ip, body)

	selector := handler.HandlerSelector{}
	go selector.Selects(&conn, message)
	reader.Close()

	handleConnection(conn)
}

func startTCPServer() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", model.TCPServerInstance.PORT))
	checkError(err)
	log.Printf("[INFO] start listening on %d", model.TCPServerInstance.PORT)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("\033[1;33m[WARNING] \033[0maccept failed", err)
			continue
		}
		log.Println("[INFO] conn# ", reflect.TypeOf(conn), conn)
		go handleConnection(conn)
	}
}

func checkError(err error) {
	if err != nil {
		log.Println("FATAL: ", err)
		os.Exit(-1)
	}
}

func startTimer(f func(int64)) {
	go func() {
		for {
			now := time.Now()
			f(utils.NowMillisecondsByTime(now))
			next := now.Add(time.Millisecond * 100)
			next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), next.Second(), next.Nanosecond(), next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}
