package handler

import (
	"log"
	"net"

	"com.lueey.shop/config"
	"com.lueey.shop/model"
	avro "com.lueey.shop/protocol"
)

type CustomerApplyJoinRoom struct {
	HandlerSelector
}

func (h *CustomerApplyJoinRoom) setConn(conn *net.Conn) {
	h.conn = conn
}

func (h *CustomerApplyJoinRoom) selected(s *model.Session) {
	h.session = s
}

func (h *CustomerApplyJoinRoom) do(msg avro.Message) {
	if model.GlobalState != avro.GlobalStateChat_with_advisor && !config.DEBUG {
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    "当前阶段不可以申请进入房间",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		h.session.SendMessage(msg)
		return
	}
	if h.session.CurrentUser() == nil {
		log.Println("ERROR: currentUser is nil, please signin first.")
		h.session.Close("session.cutomerInfo is nil.")
		return
	}

	log.Printf("DEBUG: session currentUser %s", h.session.CurrentUser())

	if h.session.CurrentUser().SignedContract {
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    "恭喜，您已签约成功",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		h.session.SendMessage(msg)
		return
	}

	r, ok := model.RoomContainer[h.session.CurrentUser().SalesAdvisorID]
	if !ok {
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    "无法进入房间，您没有受销售顾问的邀请",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		h.session.SendMessage(msg)
		return
	}

	for _, v := range r.WaitingList {
		if v == h.session.CurrentUser().ID {
			msg := *model.GenerateMessage(avro.ActionError_message)
			msg.Error_message = &avro.Error_messageUnion{
				String:    "已经在等待队伍中",
				UnionType: avro.Error_messageUnionTypeEnumString,
			}
			h.session.SendMessage(msg)
			return
		}
	}

	r.NewCustomerJoinWaitingList(h.session.CurrentUser().ID)

	roominfo := r.BuildRoomMessage()
	refreshMsg := model.GenerateMessage(avro.ActionMessage_room_info)
	refreshMsg.Message_room_info = &avro.Message_room_infoUnion{
		MessageRoomInfo: roominfo,
		UnionType:       avro.Message_room_infoUnionTypeEnumMessageRoomInfo,
	}

	// 通知销售端
	if salesAdvisorSession, exist := model.GetSessionByName(r.SalesAdvisorID); exist {
		salesAdvisorSession.SendMessage(*refreshMsg)
	}

	// 通知所有用户端
	for _, cid := range r.WaitingList {
		if waitingCustomerSession, exist := model.GetSessionByName(cid); exist {
			waitingCustomerSession.SendMessage(*refreshMsg)
		}
	}
}
