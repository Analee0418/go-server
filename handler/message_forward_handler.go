package handler

import (
	"encoding/json"
	"log"
	"net"

	"com.lueey.shop/model"
	avro "com.lueey.shop/protocol"
	"com.lueey.shop/utils"
)

// MessageForwardHandler 消息转发

type MessageForwardHandler struct {
	HandlerSelector
}

func (h *MessageForwardHandler) setConn(conn *net.Conn) {
	h.conn = conn
}

func (h *MessageForwardHandler) selected(s *model.Session) {
	h.session = s
}

func (h *MessageForwardHandler) do(msg avro.Message) {
	switch msg.Action {
	case avro.ActionMessage_forward_to_customer:
		if h.session.Room() == nil {
			log.Println("\033[1;31m[ERROR] \033[0msales roomInfo is nil, please signin first.")
			h.session.Close("session.roomInfo is nil.")
			return
		}
		r := h.session.Room()
		if r.CurrentCustomerID == "" {
			msg := *model.GenerateMessage(avro.ActionError_message)
			msg.Error_message = &avro.Error_messageUnion{
				String:    "当前没有正在签约的客户",
				UnionType: avro.Error_messageUnionTypeEnumString,
			}
			h.session.SendMessage(msg)
			return
		}
		customerSession, exist := model.GetSessionByName(r.CurrentCustomerID)
		if !exist {
			msg := *model.GenerateMessage(avro.ActionError_message)
			msg.Error_message = &avro.Error_messageUnion{
				String:    "用户端目前处于离线状态，请稍后再试",
				UnionType: avro.Error_messageUnionTypeEnumString,
			}
			h.session.SendMessage(msg)
			return
		}
		customerSession.SendMessage(msg)
	case avro.ActionMessage_forward_to_sales_advisor:
		if h.session.CurrentUser() == nil {
			log.Println("\033[1;31m[ERROR] \033[0mcurrentUser is nil, please signin first.")
			h.session.Close("session.cutomerInfo is nil.")
			return
		}
		r, ok := model.RoomContainer[h.session.CurrentUser().SalesAdvisorID]
		if !ok {
			msg := *model.GenerateMessage(avro.ActionError_message)
			msg.Error_message = &avro.Error_messageUnion{
				String:    "没有找到销售顾问对您发起邀请的记录",
				UnionType: avro.Error_messageUnionTypeEnumString,
			}
			h.session.SendMessage(msg)
			return
		}
		if r.CurrentCustomerID != h.session.CurrentUser().ID {
			msg := *model.GenerateMessage(avro.ActionError_message)
			msg.Error_message = &avro.Error_messageUnion{
				String:    "您当前并没有在房间内",
				UnionType: avro.Error_messageUnionTypeEnumString,
			}
			h.session.SendMessage(msg)
			return
		}
		salesAdvisorSession, exist := model.GetSessionByName(r.SalesAdvisorID)
		if !exist {
			msg := *model.GenerateMessage(avro.ActionError_message)
			msg.Error_message = &avro.Error_messageUnion{
				String:    "销售端目前处于离线状态，请稍后再试",
				UnionType: avro.Error_messageUnionTypeEnumString,
			}
			h.session.SendMessage(msg)
			return
		}
		salesAdvisorSession.SendMessage(msg)
	case avro.ActionMessage_broadcast:
		// TODO
		broadcastMessage := map[string]string{
			"key": msg.Message_broadcast.MessageForward.Key.String,
			"sec": msg.Message_broadcast.MessageForward.Sec.String,
		}

		lang, err := json.Marshal(broadcastMessage)
		if err == nil {
			pubErr := utils.PublishMessage("global_broadcast", string(lang))
			if pubErr != nil {
				log.Printf("\033[1;31m[ERROR] \033[0mpublish broadcast message failed, because %v", pubErr)
			}
			log.Printf("[INFO] publish broadcast message to global %s", string(lang))
		}
	}
}
