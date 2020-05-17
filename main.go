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
			log.Println("Fatal caught panic in main()", x)
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
		// log.Printf("Time. time %s", (time.Now().Nanosecond() / 1e6))
		// sessionca
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
				log.Println("Server program exit:", s)
				model.TCPServerInstance.TCPServerUpdateStatus(model.ServerDeactivate, false)
				utils.HDelRedis("hallserver##startup", model.TCPServerInstance.ID)
				log.Printf("Notify HTTPServer")
				os.Exit(0)
			case syscall.SIGUSR1:
				log.Println("usr1", s)
			case syscall.SIGUSR2:
				log.Println("usr2", s)
			default:
				log.Println("Other:", s)
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
			log.Println("ERROR: caught panic in handleConnection", x)
		}
		time.Sleep(10 * time.Microsecond)
	}()

	ip := conn.RemoteAddr().String()

	_ret := make([]byte, 4)
	n, e := conn.Read(_ret)
	if e != nil {
		if e == io.EOF { // 客户端主动断开
			log.Println("WARN: !!!!!!!!!!!!!!!!!!!!!!!!!!!!!! discover client disconnect.", ip)
			model.OnClientDisconnect(conn)
			return
		}
	}
	dataLen := int(binary.LittleEndian.Uint32(_ret))
	log.Println("->->->->->->->->->->->-> length", dataLen, n, ip, e)

	if dataLen >= 2097152 {
		log.Println("ERROR: Illegal request ->->->->->->->->->->->-> ", ip)
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
		log.Printf("len: %d, read len: %d, remain: %d\n", l, n, remain)
		if n == 0 {
			break
		}
		ulimitBuffer.Write(pkg[:n])

		remain = dataLen - len(ulimitBuffer.Bytes())
		if remain <= 0 {
			break
		}
	}

	log.Println("->->->->->->->->->->->-> read all length", ulimitBuffer.Len())

	// rr := bufio.NewReader(conn)
	// var buf [buffSize]byte
	// n, err := rr.Read(buf[:]) // 读取数据
	// if err != nil {
	// 	log.Println("read from client failed, err:", err)
	// 	return
	// }
	// if err != nil {
	// 	log.Println(err)
	// 	conn.Close()
	// 	return
	// }
	message := avro.Message{}
	deser, err := compiler.CompileSchemaBytes([]byte(message.Schema()), []byte(message.Schema()))
	if err != nil {
		log.Println("error", err)
		return
	}
	// log.Println("-> Request ", n, " Byte")
	buffer := bytes.NewBuffer(ulimitBuffer.Bytes())
	var reader = flate.NewReader(buffer)
	errs := vm.Eval(reader, deser, &message)
	if errs != nil {
		log.Printf("ERROR invalid data package: %v\n", errs)
		return
	}

	log.Println(message.Action, " Request ", ulimitBuffer.Len(), " Byte from ", ip)
	if config.DEBUG {
		lang, err := json.MarshalIndent(message, "", "   ")
		if err == nil {
			log.Println(string(lang))
		}
	}

	selector := handler.HandlerSelector{}
	go selector.Selects(&conn, message)
	reader.Close()

	handleConnection(conn)
}

func startTCPServer() {
	// var port string = ""
	// args := os.Args[1:]
	// if len(args) == 0 {
	// 	port = pORT
	// } else {
	// 	port = args[0]
	// }
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", model.TCPServerInstance.PORT))
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	checkError(err)
	log.Printf("start listening on %d", model.TCPServerInstance.PORT)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept failed", err)
			continue
		}
		log.Println("conn: ", reflect.TypeOf(conn), conn)
		go handleConnection(conn)
	}
}

func checkError(err error) {
	if err != nil {
		log.Println("Fatal error:", err)
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
