package main

import (
	"bufio"
	"bytes"
	"compress/flate"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"time"

	"com.lueey.shop/handler"
	"com.lueey.shop/model"
	avro "com.lueey.shop/protocol"
	"com.lueey.shop/utils"
	"github.com/actgardner/gogen-avro/compiler"
	"github.com/actgardner/gogen-avro/vm"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
}

func main() {
	defer func() {
		if x := recover(); x != nil {
			log.Println("caught panic in main()", x)
		}
	}()

	runtime.GOMAXPROCS(runtime.NumCPU())

	startTimer(func(now int64) {
		// log.Printf("Time. time %s", (time.Now().Nanosecond() / 1e6))
		// sessionca
		model.ReleaseSessionCache(now)
	})
	startTCPServer()
}

const (
	buffSize int    = 1024 * 10
	pORT     string = "1234"
)

func handleConnection(conn net.Conn) {
	defer func() {
		if x := recover(); x != nil {
			log.Println("caught panic in handleConnection", x)
		}
		time.Sleep(10 * time.Microsecond)
	}()

	rr := bufio.NewReader(conn)
	var buf [buffSize]byte
	n, err := rr.Read(buf[:]) // 读取数据
	if err != nil {
		fmt.Println("read from client failed, err:", err)
		return
	}
	if err != nil {
		log.Println(err)
		conn.Close()
		return
	}
	message := avro.Message{}
	deser, err := compiler.CompileSchemaBytes([]byte(message.Schema()), []byte(message.Schema()))
	if err != nil {
		log.Println("error", err)
		return
	}
	log.Println("Recive ", n, " Byte")
	buffer := bytes.NewBuffer(buf[:n])
	var reader = flate.NewReader(buffer)
	errs := vm.Eval(reader, deser, &message)
	if errs != nil {
		log.Printf("ERROR invalid data package: %v\n", errs)
		return
	}

	log.Println(message)
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
