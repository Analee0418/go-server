package handler

import (
	"log"
	"net"

	"com.lueey.shop/model"
	avro "com.lueey.shop/protocol"
)

type CustomerSignin struct {
	HandlerSelector
}

func (h *CustomerSignin) setConn(conn *net.Conn) {
	h.conn = conn
}

func (h *CustomerSignin) selected(s *model.Session) {
	h.session = s
}

func (h *CustomerSignin) do(msg avro.Message) {
	Idcard := msg.Customer_signin.RequestCustomerSignin.Idcard.String
	mobile := msg.Customer_signin.RequestCustomerSignin.Mobile.String

	if s, exists := model.GetSessionByName(Idcard); exists {
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    "Do not log in repeatedly",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		s.SendMessage(msg)
		return
	}

	ok := false
	keys := []string{}
	for k, _ := range model.AllCustomerContainer {
		keys = append(keys, k)
	}
	log.Println("-----------------------", keys, Idcard)
	c, ok := model.AllCustomerContainer[Idcard]
	if !ok {
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    "Can not found user",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		model.SendMessage(*h.conn, msg)
		return
	}
	if c.CustomerInfo.Mobile.String == mobile {
		ok = true
	}

	// 允许登录
	if ok {
		h.session = new(model.Session)
		h.session.InitCustomer(*h.conn, Idcard)

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

		keys = []string{}
		for k, _ := range model.RoomContainer {
			keys = append(keys, k)
		}

		log.Println(keys, c.SalesAdvisorID)

		if r, ok := model.RoomContainer[model.GenerateRoomKey(c.SalesAdvisorID)]; ok {
			log.Printf("Sales advisor room: %v", r)

			roominfo := r.GetRoomInfo()
			msg := *model.GenerateMessage(avro.ActionMessage_room_info)
			msg.Message_room_info = &avro.Message_room_infoUnion{
				MessageRoomInfo: &roominfo,
				UnionType:       avro.Message_room_infoUnionTypeEnumMessageRoomInfo,
			}
			h.session.SendMessage(msg)
		} else {
			msg := *model.GenerateMessage(avro.ActionError_message)
			msg.Error_message = &avro.Error_messageUnion{
				String:    "Invalid users",
				UnionType: avro.Error_messageUnionTypeEnumString,
			}
			model.SendMessage(*h.conn, msg)
		}
	}
}
