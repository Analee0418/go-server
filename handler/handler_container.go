package handler

import (
	"log"
	"net"

	"com.lueey.shop/model"
	avro "com.lueey.shop/protocol"
)

var handlerMapping = map[avro.Action]ActionHandler{
	avro.ActionHeartbeat:                    &heartBeat{},
	avro.ActionRequest_sales_advisor_signin: &SalesAdvisorSignin{},
	avro.ActionRequest_customer_signin:      &CustomerSignin{},
}

type ActionHandler interface {
	do(msg avro.Message)
	setConn(conn *net.Conn)
	selected(session *model.Session)
}

type HandlerSelector struct {
	name    string
	conn    *net.Conn
	session *model.Session
}

func (s *HandlerSelector) Selects(conn *net.Conn, msg avro.Message) {
	handler, ok := handlerMapping[msg.Action]
	if ok {
		s.conn = conn
		handler.setConn(conn)
		handler.do(msg)
		cacheSession, exist := model.SessionByID(msg.SessionId.String)
		if exist {
			s.session = cacheSession
			handler.selected(cacheSession)
			log.Printf("exist: %v, cacheSession: %v", exist, cacheSession)
		}
		handler.do(msg)
	} else {
		log.Printf("Action not found: %v", msg.Action)
	}
}

type heartBeat struct {
	HandlerSelector
}

func (h *heartBeat) setConn(conn *net.Conn) {
	h.conn = conn
}

func (h *heartBeat) selected(s *model.Session) {
	h.session = s
}

func (h *heartBeat) do(msg avro.Message) {
	_conn := *h.conn
	log.Printf("[%v] heartbeat message %v\n", _conn.RemoteAddr().String(), h.session)
	if h.session == nil {
		log.Println("login first pls")
		msg := model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    "login first pls",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		model.SendMessage(_conn, *msg)
	} else {
		h.session.Heartbeat()
		h.session.SendMessage(*model.GenerateMessage(avro.ActionHeartbeat))
		log.Printf("%v", h.session)
	}
}
