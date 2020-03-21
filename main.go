package main

import (
	"bufio"
	"bytes"
	"compress/flate"
	"fmt"
	"log"
	"net"
	"os"

	"com.lueey.shop/handler"
	avro "com.lueey.shop/protocol"
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

	startTCPServer()
}

const (
	HEAD_SIZE = 8
	BUF_SIZE  = 1048576
)

func handleConnection(conn net.Conn) {
	defer func() {
		if x := recover(); x != nil {
			log.Println("caught panic in handleConnection", x)
		}
	}()

	rr := bufio.NewReader(conn)
	var buf [BUF_SIZE]byte
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
		log.Println(errs)
	}

	log.Println(message)
	selector := handler.HandlerSelector{}
	go selector.Selects(&conn, message)
	reader.Close()

	handleConnection(conn)
}

func startTCPServer() {
	listener, err := net.Listen("tcp", ":1234")
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	checkError(err)
	log.Println("start listening on 1234")
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
