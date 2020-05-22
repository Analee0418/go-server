package handler

import (
	"fmt"
	"log"
	"net"

	"com.lueey.shop/config"
	"com.lueey.shop/model"
	avro "com.lueey.shop/protocol"
)

type CustomerUpdateStateHanlder struct {
	HandlerSelector
}

func (h *CustomerUpdateStateHanlder) setConn(conn *net.Conn) {
	h.conn = conn
}

func (h *CustomerUpdateStateHanlder) selected(s *model.Session) {
	h.session = s
}

func (h *CustomerUpdateStateHanlder) do(msg avro.Message) {
	if h.session.CurrentUser() == nil {
		log.Printf("\033[1;31m[ERROR] \033[0mcurrentUser is nil, please signin first. session: %s", h.session)
		h.session.Close("session.cutomerInfo is nil.")
		return
	}

	if config.DEBUG {
		log.Printf("[DEBUG] session currentUser %s", h.session.CurrentUser())
	}

	if h.session.CurrentUser().SignedContract {
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    "恭喜，您已签约成功",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		h.session.SendMessage(msg)
		return
	}

	_, ok := model.RoomContainer[h.session.CurrentUser().SalesAdvisorID]
	if !ok {
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    "无法更新状态，您没有受销售顾问的邀请",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		h.session.SendMessage(msg)
		return
	}

	if h.session.CurrentUser().State == string(msg.Customer_update_state.RequestCustomerUpdateState.State) {
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    fmt.Sprintf("当前已经是该状态[%s]，无需改变", h.session.CurrentUser().State),
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		h.session.SendMessage(msg)
		return
	}

	// 更新状态
	h.session.CurrentUser().ChangeState(msg.Customer_update_state.RequestCustomerUpdateState.State)
	// 刷新前端
	msg = *model.GenerateMessage(avro.ActionMessage_customers_info)
	msg.Message_customer_info = &avro.Message_customer_infoUnion{
		MessageCustomersInfo: h.session.CurrentUser().BuildCustomerMessage(),
		UnionType:            avro.Message_customer_infoUnionTypeEnumMessageCustomersInfo,
	}
	h.session.SendMessage(msg)
}
