package handler

import (
	"log"
	"net"
	"runtime/debug"

	"com.lueey.shop/model"
	avro "com.lueey.shop/protocol"
)

var handlerMapping = map[avro.Action]ActionHandler{
	avro.ActionHeartbeat:                                 &heartBeat{},
	avro.ActionRequest_sales_advisor_signin:              &SalesAdvisorSignin{},
	avro.ActionRequest_sales_advisor_receiving_customers: &SalesAdvisorInvite{},
	avro.ActionRequest_sales_advisor_leave_customers:     &SalesAdvisorKick{},
	avro.ActionRequest_sales_advisor_build_contract:      &SalesAdvisorBuildContract{},
	avro.ActionRequest_sales_advisor_confirm_paid:        &SalesAdvisorConfirmedSignedContract{},
	avro.ActionMessage_cars_model:                        &SalesAdvisorUpdateCarModel{},
	avro.ActionRequest_customer_signin:                   &CustomerSignin{},
	avro.ActionRequest_customer_join_queue:               &CustomerApplyJoinRoom{},
	avro.ActionRequest_customer_build_signature:          &CustomerBuildSignature{},
	avro.ActionMessage_forward_to_customer:               &MessageForwardHandler{},
	avro.ActionMessage_forward_to_sales_advisor:          &MessageForwardHandler{},
	avro.ActionMessage_broadcast:                         &MessageForwardHandler{},
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
	defer func() {
		if x := recover(); x != nil {
			log.Println("\033[1;31m[ERROR] \033[0mcaught panic in handleConnection", x)
			debug.PrintStack()
		}
	}()

	handler, ok := handlerMapping[msg.Action]
	if ok {
		s.conn = conn
		handler.setConn(conn)

		// 登录不需要session外，其他请求都需要sesison
		if _, ok := map[avro.Action]string{
			avro.ActionRequest_customer_signin:      "cutomer_signin",
			avro.ActionRequest_sales_advisor_signin: "salce_signin",
			// avro.ActionRequest_sales_advisor_signin: "salce_signin",
		}[msg.Action]; ok { // 登录操作

			handler.do(msg)

		} else { // 其他行为都需要提前具备 session 实例

			cacheSession, exist := model.SessionByID(msg.SessionId.String)
			if exist && !cacheSession.Dead() {
				s.session = cacheSession
				handler.selected(cacheSession)
				log.Printf("[INFO] exist: %v, cacheSession: %v", exist, cacheSession)
				handler.do(msg)
			} else {
				log.Println("\033[1;31m[ERROR] \033[0mlogin first pls")
				msg := model.GenerateMessage(avro.ActionError_message)
				msg.Error_message = &avro.Error_messageUnion{
					String:    "请先登录账户",
					UnionType: avro.Error_messageUnionTypeEnumString,
				}
				model.SendMessage(*s.conn, *msg)
				return
			}
		}
	} else {
		log.Printf("\033[1;31m[ERROR] \033[0mAction not found: %v", msg.Action)
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
	log.Printf("[INFO] [%v] heartbeat message %v\n", _conn.RemoteAddr().String(), h.session)
	h.session.Heartbeat()
	h.session.SendMessage(*model.GenerateMessage(avro.ActionHeartbeat))
	log.Printf("[INFO] %s", h.session)
}
