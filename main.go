package main

import (
	"bytes"
	"compress/flate"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"time"

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
	// defer func() {
	// 	if x := recover(); x != nil {
	// 		log.Println("caught panic in main()", x)
	// 	}
	// }()

	runtime.GOMAXPROCS(runtime.NumCPU())

	config.Init()

	model.InitAuctionGoods()
	model.InitCustomer()
	model.InitRoom()
	model.InitGlobal()

	model.PostInitAuctionGoods()
	model.PostInitCustoemr()
	model.PostInitRoom()
	model.PostInitGlobal()

	startTimer(func(now int64) {
		// log.Printf("Time. time %s", (time.Now().Nanosecond() / 1e6))
		// sessionca
		model.ReleaseSessionCache(now)
	})

	go utils.ReceiveGlobalState()

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

	_ret := make([]byte, 4)
	n, e := conn.Read(_ret)
	dataLen := int(binary.LittleEndian.Uint32(_ret))
	log.Println("->->->->->->->->->->->-> length", dataLen, e)

	ulimitBuffer := bytes.NewBuffer(nil)
	remain := dataLen
	for {
		l := buffSize
		if l < dataLen {
			if l > remain {
				l = remain
			}
		}
		pkg := make([]byte, l)
		n, _ := conn.Read(pkg)
		log.Printf("len: %d, read len: %d, remain: %d\n", l, n, remain)
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

	// lang, err := json.Marshal(message)
	// if err == nil {
	// 	log.Println(string(lang))
	// }
	log.Println(message.Action, " Request ", n, " Byte")

	selector := handler.HandlerSelector{}
	go selector.Selects(&conn, message)
	reader.Close()

	handleConnection(conn)
}

func startTCPServer() {
	var port string = ""
	args := os.Args[1:]
	if len(args) == 0 {
		port = pORT
	} else {
		port = args[0]
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	checkError(err)
	log.Printf("start listening on %s", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept failed", err)
			continue
		}
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
			next := now.Add(time.Second * 10)
			next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), next.Second(), next.Nanosecond(), next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}
