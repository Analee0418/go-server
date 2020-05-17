package handler

import (
	"log"
	"net"

	"com.lueey.shop/config"
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

	if _, exists := model.GetSessionByConn(*h.conn); exists {
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    "该设备已经登录了其他账号",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		model.SendMessage(*h.conn, msg)
		return
	}

	log.Printf("DEBUG customer signin with %s, %s", Idcard, mobile)

	var ok = false
	var customer *model.Customer = nil

	customer, ok = model.AllCustomerContainer[Idcard]
	if !ok || customer == nil {
		for _, c := range model.AllCustomerContainer {
			if l := len(c.ID); c.ID[l-6:l] == Idcard && c.Mobile == mobile {
				Idcard = c.ID
				customer = c
				break
			}
		}
	}

	if _, exists := model.GetSessionByName(Idcard); exists {
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    "您已在其他设备登录",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		model.SendMessage(*h.conn, msg)
		return
	}

	if customer == nil || customer.Mobile != mobile {
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    "无效用户",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		model.SendMessage(*h.conn, msg)
		return
	}

	_, ok = model.RoomContainer[customer.SalesAdvisorID]
	if !ok {
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    "您没有受销售顾问的邀请",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		model.SendMessage(*h.conn, msg)
		return
	}

	// 允许登录
	if ok {
		h.session = new(model.Session)
		h.session.InitCustomer(*h.conn, customer)

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
		if config.DEBUG {
			log.Println("DEBUG: On user login successed!", smsg.Message_session.MessageSession.Sid.String)
		}
		h.session.SendMessage(*smsg)

		model.TCPServerInstance.TCPServerOnUpdateOnlines(model.Onlines(), false)

		msgs := model.TCPGlobalOnCustomerSignin(Idcard)
		if msgs != nil {

			for _, msg := range msgs {
				h.session.SendMessage(*msg)
			}

			// if r, ok := model.RoomContainer[model.GenerateRoomKey(c.SalesAdvisorID)]; ok {
			// 	log.Printf("Sales advisor room: %v", r)

			// 	roominfo := r.BuildRoomMessage()
			// 	msg := *model.GenerateMessage(avro.ActionMessage_room_info)
			// 	msg.Message_room_info = &avro.Message_room_infoUnion{
			// 		MessageRoomInfo: roominfo,
			// 		UnionType:       avro.Message_room_infoUnionTypeEnumMessageRoomInfo,
			// 	}
			// 	h.session.SendMessage(msg)
			// } else {
			// 	msg := *model.GenerateMessage(avro.ActionError_message)
			// 	msg.Error_message = &avro.Error_messageUnion{
			// 		String:    "Invalid users",
			// 		UnionType: avro.Error_messageUnionTypeEnumString,
			// 	}
			// 	model.SendMessage(*h.conn, msg)
			// }
		}
	}
}
