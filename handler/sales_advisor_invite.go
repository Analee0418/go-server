package handler

import (
	"fmt"
	"log"
	"net"

	"com.lueey.shop/model"
	avro "com.lueey.shop/protocol"
)

// SalesAdvisorInvite 邀请客户进入房间

type SalesAdvisorInvite struct {
	HandlerSelector
}

func (h *SalesAdvisorInvite) setConn(conn *net.Conn) {
	h.conn = conn
}

func (h *SalesAdvisorInvite) selected(s *model.Session) {
	h.session = s
}

func (h *SalesAdvisorInvite) do(msg avro.Message) {
	if h.session.Room() == nil {
		log.Println("\033[1;31m[ERROR] \033[0msales roomInfo is nil, please signin first.")
		h.session.Close("session.roomInfo is nil.")
		return
	}

	r := h.session.Room()

	if r.CurrentCustomerID != "" {
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    fmt.Sprintf("目前正在与 %s 洽谈中", r.CurrentCustomerID),
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		h.session.SendMessage(msg)
		return
	}

	idcard := msg.Sales_advisor_receiving_customers.RequestSalesAdvisorReceivingCustomers.Customers_idcard

	var c *model.Customer = nil
	for _, cid := range r.WaitingList {
		if cid == idcard {
			if customer, ok := model.AllCustomerContainer[cid]; ok {
				c = customer
			}
			break
		}
	}

	if c == nil {
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    "没有在等待队列中找到该用户",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		h.session.SendMessage(msg)
		return
	}

	if c.SignedContract {
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    "该用户已经签约成功",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		h.session.SendMessage(msg)
		return
	}

	if c.SalesAdvisorID != r.SalesAdvisorID {
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    "用户数据 SalesAdvisorID 错误，请联系服务器",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		h.session.SendMessage(msg)
		return
	}

	r.DeleteWaiting(idcard)
	r.UpdateCustomer(idcard)

	log.Printf("\033[1;36mSTATS: \033[0m%s Invite %s join room to chatting", r.SalesAdvisorID, idcard)

	roominfo := r.BuildRoomMessage()
	refreshMsg := model.GenerateMessage(avro.ActionMessage_room_info)
	refreshMsg.Message_room_info = &avro.Message_room_infoUnion{
		MessageRoomInfo: roominfo,
		UnionType:       avro.Message_room_infoUnionTypeEnumMessageRoomInfo,
	}

	// 刷新销售端
	h.session.SendMessage(*refreshMsg)

	// 刷新所有等待的用户端
	for _, cid := range r.WaitingList {
		if waitingCustomerSession, exist := model.GetSessionByName(cid); exist {
			waitingCustomerSession.SendMessage(*refreshMsg)
		}
	}

	// 刷新当事人用户端
	if waitingCustomerSession, exist := model.GetSessionByName(idcard); exist {
		waitingCustomerSession.SendMessage(*refreshMsg)

		invitedMsg := model.GenerateMessage(avro.ActionMessage_just_been_invited_into_room)
		waitingCustomerSession.SendMessage(*invitedMsg)
	}
}

// SalesAdvisorKick 请出客户
type SalesAdvisorKick struct {
	HandlerSelector
}

func (h *SalesAdvisorKick) setConn(conn *net.Conn) {
	h.conn = conn
}

func (h *SalesAdvisorKick) selected(s *model.Session) {
	h.session = s
}

func (h *SalesAdvisorKick) do(msg avro.Message) {
	if h.session.Room() == nil {
		log.Println("\033[1;31m[ERROR] \033[0msales roomInfo is nil, please signin first.")
		h.session.Close("session.roomInfo is nil.")
		return
	}

	r := h.session.Room()

	old := r.CurrentCustomerID
	if r.CurrentCustomerID == "" {
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    "目前洽谈室没人",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		h.session.SendMessage(msg)
		return
	}

	r.UpdateCustomer("")

	roominfo := r.BuildRoomMessage()
	refreshMsg := model.GenerateMessage(avro.ActionMessage_room_info)
	refreshMsg.Message_room_info = &avro.Message_room_infoUnion{
		MessageRoomInfo: roominfo,
		UnionType:       avro.Message_room_infoUnionTypeEnumMessageRoomInfo,
	}

	// 刷新销售端
	h.session.SendMessage(*refreshMsg)

	// 等值当事人用户端聊天结束
	if waitingCustomerSession, exist := model.GetSessionByName(old); exist {
		evenMessage := model.GenerateMessage(avro.ActionMessage_room_chat_ends)
		waitingCustomerSession.SendMessage(*evenMessage)
	}
}
