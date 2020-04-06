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

func (h *SalesAdvisorSignin) selected(s *model.Session) {
	h.session = s
}

func (h *SalesAdvisorSignin) do(msg avro.Message) {
	// alias := msg.Sales_advisor_signin.RequestSalesAdvisorSignin.Sales_advisor_alias.String
	advisorID := msg.Sales_advisor_signin.RequestSalesAdvisorSignin.Sales_advisor_id.String

	log.Printf("DEBUG sales advisor login on %s", advisorID)

	if _, exists := model.GetSessionByConn(*h.conn); exists {
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    "该设备已经登录了其他账号",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		model.SendMessage(*h.conn, msg)
		return
	}

	if _, exists := model.GetSessionByName(advisorID); exists {
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    "已在其他设备上登录",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		model.SendMessage(*h.conn, msg)
		return
	}

	// 允许登录
	if r, ok := model.RoomContainer[advisorID]; ok {
		log.Printf("find room %s", advisorID)
		log.Printf("Sales advisor room: %v", r)
		h.session = new(model.Session)
		h.session.InitAdvisor(*h.conn, r)

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

		roominfo := r.BuildRoomMessage()
		log.Println(r)
		msg := *model.GenerateMessage(avro.ActionMessage_room_info)
		msg.Message_room_info = &avro.Message_room_infoUnion{
			MessageRoomInfo: roominfo,
			UnionType:       avro.Message_room_infoUnionTypeEnumMessageRoomInfo,
		}

		h.session.SendMessage(msg)
	}
}
