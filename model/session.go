package model

import (
	"bytes"
	"compress/flate"
	"fmt"
	"log"
	"net"

	avro "com.lueey.shop/protocol"
	guuid "github.com/google/uuid"
)

type Session struct {
	conn         net.Conn
	id           guuid.UUID
	ip           string
	customerInfo avro.MessageCustomersInfo
	salesAdvisor string
}

func (s *Session) String() string {
	return fmt.Sprintf("%v/%v", s.ip, s.id.String())
}

func (s *Session) InitAdvisor(conn net.Conn, salesAdvisor string) {
	s.conn = conn
	s.id = guuid.New()
	s.ip = s.conn.RemoteAddr().String()
	s.salesAdvisor = salesAdvisor
	AddSession(s)
}

func (s *Session) InitCustomer(conn net.Conn, customer avro.MessageCustomersInfo) {
	s.conn = conn
	s.id = guuid.New()
	AddSession(s)
}

func SendMessage(conn net.Conn, msg avro.Message) {
	blockBytes := make([]byte, 0)
	blockBuffer := bytes.NewBuffer(blockBytes)

	compressedWriter, err := flate.NewWriter(blockBuffer, flate.DefaultCompression)
	if err != nil {
		log.Println(err)
	}
	msg.Serialize(compressedWriter)
	compressedWriter.Flush()

	log.Println(blockBuffer.Len())
	defer func() {
		compressedWriter.Close()
	}()

	// head := make([]byte, HEAD_SIZE)
	content := blockBuffer.Bytes()
	// headSize := blockBuffer.Len()
	// binary.BigEndian.PutUint16(head, uint16(headSize))

	//先写入head部分，再写入body部分
	// _, err = conn.Write(head)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return err
	// }
	_, err = conn.Write(content)
	if err != nil {
		log.Println(err)
	}
}

func (s *Session) SendMessage(message avro.Message) {
	SendMessage(s.conn, message)
}

func (s *Session) close() {
	if s == nil {
		log.Print()
		return
	}
	msg := avro.Message{
		Action:            avro.ActionMessage_room_info,
		Message_room_info: avro.NewUnionNullMessageRoomInfo(),
	}
	s.SendMessage(msg)
	s.conn.Close()
	DeleteSession(s)
}
