package model

import (
	"bytes"
	"compress/flate"
	"fmt"
	"log"
	"net"
	"time"

	avro "com.lueey.shop/protocol"
	"com.lueey.shop/utils"
	guuid "github.com/google/uuid"
)

type Session struct {
	conn                     net.Conn
	id                       guuid.UUID
	ip                       string
	customerInfo             avro.MessageCustomersInfo
	name                     string
	salesAdvisor             string
	lastHeartBeat            time.Time
	lastHeartBeatMillisecond int64
	dead                     bool
}

func (s *Session) UUID() string {
	return s.id.String()
}

func (s *Session) String() string {
	return fmt.Sprintf("%v/%v/%v/%v/%v", s.ip, s.id.String(), s.name, s.lastHeartBeat, s.lastHeartBeatMillisecond)
}

func (s *Session) InitAdvisor(conn net.Conn, salesAdvisor string) {
	s.conn = conn
	s.id = guuid.New()
	s.ip = s.conn.RemoteAddr().String()
	s.name = salesAdvisor
	s.salesAdvisor = salesAdvisor
	s.lastHeartBeat = time.Now()
	s.lastHeartBeatMillisecond = utils.NowMillisecondsByTime(s.lastHeartBeat)
	AddSession(s)
}

func (s *Session) InitCustomer(conn net.Conn, customer avro.MessageCustomersInfo) {
	s.conn = conn
	s.id = guuid.New()
	AddSession(s)
	// TODO
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
	if s.dead {
		log.Printf("Warn, the sesison[%v, %s] has closed\n", s.id, s.name)
		return
	}
	SendMessage(s.conn, message)
}

func (s *Session) Heartbeat() {
	now := time.Now()
	s.lastHeartBeat = now
	s.lastHeartBeatMillisecond = utils.NowMillisecondsByTime(now)
}

func (s *Session) Close(reason string) (guuid.UUID, string) {
	msg := GenerateMessage(avro.ActionError_message)
	msg.Error_message = &avro.Error_messageUnion{String: reason, UnionType: avro.Error_messageUnionTypeEnumString}

	s.SendMessage(*msg)
	s.conn.Close()
	s.dead = true

	// 如果在竞拍
	// 如果在排队
	// 如果已进入房间
	//

	return s.id, s.name
}
