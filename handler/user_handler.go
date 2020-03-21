package handler

import (
	"log"
	"net"

	"com.lueey.shop/model"
	avro "com.lueey.shop/protocol"
)

type SalesAdvisorSignin struct {
	HandlerSelector
}

func (h *SalesAdvisorSignin) setConn(conn *net.Conn) {
	h.conn = conn
}

func (h *SalesAdvisorSignin) do(msg avro.Message) {
	// alias := msg.Sales_advisor_signin.RequestSalesAdvisorSignin.Sales_advisor_alias.String
	advisorID := msg.Sales_advisor_signin.RequestSalesAdvisorSignin.Sales_advisor_id.String

	if s, exists := model.GetSessionByName(advisorID); exists {
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message.String = "请勿重复登录"
		s.SendMessage(msg)
		return
	}

	ok := false
	// 需要提前预定一批 顾问列表
	for _, aid := range model.AdvisorID {
		if aid == advisorID {
			ok = true
			break
		}
	}
	// 允许登录
	if ok {
		if r, ok := model.RoomContainer[model.GenerateRoomKey(advisorID)]; ok {
			log.Printf("Sales advisor room: %v", r)
			h.session = new(model.Session)
			h.session.InitAdvisor(*h.conn, advisorID)
			roominfo := r.GetRoomInfo()
			msg := *model.GenerateMessage(avro.ActionMessage_room_info)
			msg.Message_room_info = &avro.UnionNullMessageRoomInfo{
				MessageRoomInfo: &roominfo,
			}

			h.session.SendMessage(msg)
		}
	}
}
