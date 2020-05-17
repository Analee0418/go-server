package handler

import (
	"log"
	"net"

	"com.lueey.shop/model"
	avro "com.lueey.shop/protocol"
)

// SalesAdvisorConfirmedSignedContract 确认成功签约
type SalesAdvisorConfirmedSignedContract struct {
	HandlerSelector
}

func (h *SalesAdvisorConfirmedSignedContract) setConn(conn *net.Conn) {
	h.conn = conn
}

func (h *SalesAdvisorConfirmedSignedContract) selected(s *model.Session) {
	h.session = s
}

func (h *SalesAdvisorConfirmedSignedContract) do(msg avro.Message) {
	if h.session.Room() == nil {
		log.Println("ERROR: sales roomInfo is nil, please signin first.")
		h.session.Close("session.roomInfo is nil.")
		return
	}

	price := msg.Sales_advisor_confirm_paid.RequestSalesAdvisorConfirmPaid.Price
	disprice := msg.Sales_advisor_confirm_paid.RequestSalesAdvisorConfirmPaid.DisPrice
	brand := msg.Sales_advisor_confirm_paid.RequestSalesAdvisorConfirmPaid.Brand.String
	color := msg.Sales_advisor_confirm_paid.RequestSalesAdvisorConfirmPaid.Color.String
	interior := msg.Sales_advisor_confirm_paid.RequestSalesAdvisorConfirmPaid.Interior.String
	series := msg.Sales_advisor_confirm_paid.RequestSalesAdvisorConfirmPaid.Series.String

	r := h.session.Room()

	old := r.CurrentCustomerID
	if r.CurrentCustomerID == "" {
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    "当前没有正在签约的客户",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		h.session.SendMessage(msg)
		return
	}

	if c, ok := model.AllCustomerContainer[r.CurrentCustomerID]; ok {
		// 标记签约成功
		c.ConfirmedSignContract(c.SalesAdvisorID, price, disprice, brand, color, interior, series)
		log.Printf("WARN: Customer[%s] signed conract ok, and confirmed payment info. The contract: price: %f, disprice: %f, brand: %s, color: %s, interior: %s, series: %s", r.CurrentCustomerID, price, disprice, brand, color, interior, series)

		// 从房间请出
		r.UpdateCustomer("")
		// 增加成交数量
		r.UpdateOrderCount(r.OrderCount + 1)
	}

	// TODO 广播签约成功

	roominfo := r.BuildRoomMessage()
	refreshMsg := model.GenerateMessage(avro.ActionMessage_room_info)
	refreshMsg.Message_room_info = &avro.Message_room_infoUnion{
		MessageRoomInfo: roominfo,
		UnionType:       avro.Message_room_infoUnionTypeEnumMessageRoomInfo,
	}

	// 刷新销售端
	h.session.SendMessage(*refreshMsg)

	// 通知当事人用户端聊天结束
	if waitingCustomerSession, exist := model.GetSessionByName(old); exist {
		evenMessage := model.GenerateMessage(avro.ActionMessage_room_chat_ends)
		waitingCustomerSession.SendMessage(*evenMessage)
	}

}
