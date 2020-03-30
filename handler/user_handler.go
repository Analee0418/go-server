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
		msg.Error_message = &avro.Error_messageUnion{
			String:    "Do not log in repeatedly",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
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

			smsg := model.GenerateMessage(avro.ActionMessage_session)
			smsg.Message_session = &avro.Message_sessionUnion{
				MessageSession: &avro.MessageSession{
					Sid: &avro.SidUnion{
						String:    h.session.UUID(),
						UnionType: avro.SidUnionTypeEnumString,
					},
				},
				UnionType: avro.Message_sessionUnionTypeEnumMessageSession,
			}
			log.Println("-----------------------", smsg.Message_session.MessageSession.Sid.String)
			h.session.SendMessage(*smsg)

			roominfo := r.GetRoomInfo()
			msg := *model.GenerateMessage(avro.ActionMessage_room_info)
			msg.Message_room_info = &avro.Message_room_infoUnion{
				MessageRoomInfo: &roominfo,
				UnionType:       avro.Message_room_infoUnionTypeEnumMessageRoomInfo,
			}

			h.session.SendMessage(msg)
		}
	}
}
