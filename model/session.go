package model

import (
	"bytes"
	"compress/flate"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"

	"com.lueey.shop/config"
	avro "com.lueey.shop/protocol"
	"com.lueey.shop/utils"
	guuid "github.com/google/uuid"
)

type Session struct {
	conn                     net.Conn
	id                       guuid.UUID
	ip                       string
	customerInfo             *Customer
	roomInfo                 *Room
	name                     string
	salesAdvisor             string
	lastHeartBeat            time.Time
	lastHeartBeatMillisecond int64
	dead                     bool
}

func (s *Session) UUID() string {
	return s.id.String()
}

func (s *Session) SetDead() {
	s.dead = true
}

func (s *Session) Dead() bool {
	// log.Printf("Session[%s] isDead: %v", s.name, s.dead)
	return s.dead
}

func (s *Session) String() string {
	return fmt.Sprintf("%v/%v/%v/%v/%v", s.ip, s.id.String(), s.name, s.lastHeartBeat, s.lastHeartBeatMillisecond)
}

func (s *Session) CurrentUser() *Customer {
	return s.customerInfo
}

func (s *Session) Room() *Room {
	return s.roomInfo
}

func (s *Session) InitAdvisor(conn net.Conn, room *Room) {
	s.conn = conn
	s.id = guuid.New()
	s.ip = s.conn.RemoteAddr().String()
	s.name = room.SalesAdvisorID
	s.salesAdvisor = room.SalesAdvisorID
	s.roomInfo = room
	s.lastHeartBeat = time.Now()
	s.lastHeartBeatMillisecond = utils.NowMillisecondsByTime(s.lastHeartBeat)
	AddSession(conn, s)
}

func (s *Session) InitCustomer(conn net.Conn, customer *Customer) {
	s.conn = conn
	s.id = guuid.New()
	s.ip = s.conn.RemoteAddr().String()
	s.name = customer.ID
	s.customerInfo = customer
	s.lastHeartBeat = time.Now()
	s.lastHeartBeatMillisecond = utils.NowMillisecondsByTime(s.lastHeartBeat)
	AddSession(conn, s)
	// TODO
}

func SendMessage(conn net.Conn, msg avro.Message) {
	blockBytes := make([]byte, 0)
	blockBuffer := bytes.NewBuffer(blockBytes)

	compressedWriter, err := flate.NewWriter(blockBuffer, flate.DefaultCompression)
	if err != nil {
		log.Println("\033[1;31mERROR: \033[0m", err)
	}
	msg.Serialize(compressedWriter)
	compressedWriter.Flush()

	defer func() {
		compressedWriter.Close()
	}()

	// head := make([]byte, HEAD_SIZE)

	lenBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(lenBytes, uint32(blockBuffer.Len()))
	if config.DEBUG {
		log.Printf("DEBUG: will write to %d, %v byte", blockBuffer.Len(), lenBytes)
	}

	finalByteArray := make([]byte, 0)
	finalBlockBuffer := bytes.NewBuffer(finalByteArray)

	finalBlockBuffer.Write(lenBytes)
	finalBlockBuffer.Write(blockBuffer.Bytes())

	// headSize := blockBuffer.Len()
	// binary.BigEndian.PutUint16(head, uint16(headSize))

	//先写入head部分，再写入body部分
	// _, err = conn.Write(head)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return err
	// }
	ip := conn.RemoteAddr().String()
	_, err = conn.Write(finalBlockBuffer.Bytes())

	if config.DEBUG {
		lang, err := json.MarshalIndent(msg, "", "   ")
		if err == nil {
			log.Printf("DEBUG: send msg to %s,\n%s", ip, string(lang))
		}
	}
	if err != nil {
		log.Printf("\033[1;31mERROR: \033[0m%s, IP: %s", err, ip)
	}
}

func (s *Session) SendMessage(message avro.Message) {
	if s.dead {
		log.Printf("\033[1;33mWARNING: \033[0mthe sesison[%v, %s] has closed\n", s.id, s.name)
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
	defer func() { s.conn.Close() }()

	msg := GenerateMessage(avro.ActionError_message)
	msg.Error_message = &avro.Error_messageUnion{String: reason, UnionType: avro.Error_messageUnionTypeEnumString}

	log.Printf("\033[1;33mWARNING: \033[0msession[%v, %s, %s] Disconnected. Reason: %s", s.id, s.name, s.customerInfo, reason)

	s.SendMessage(*msg)

	if s.customerInfo != nil {
		s.customerInfo = nil
	}

	// 如果在竞拍
	// 如果在排队
	// 如果已进入房间
	//

	s.SetDead()

	return s.id, s.name
}
