package handler

import (
	"log"
	"net"

	"com.lueey.shop/model"
	avro "com.lueey.shop/protocol"
)

// SalesAdvisorUpdateCarModel 更新汽车模型
type SalesAdvisorUpdateCarModel struct {
	HandlerSelector
}

func (h *SalesAdvisorUpdateCarModel) setConn(conn *net.Conn) {
	h.conn = conn
}

func (h *SalesAdvisorUpdateCarModel) selected(s *model.Session) {
	h.session = s
}

func (h *SalesAdvisorUpdateCarModel) do(msg avro.Message) {
	if h.session.Room() == nil {
		log.Println("\033[1;31mERROR: \033[0msales roomInfo is nil, please signin first.")
		h.session.Close("session.roomInfo is nil.")
		return
	}

	r := h.session.Room()

	r.UpdateCarModel(msg.Message_cars_model.MessageCarsModel)

	// 通知当事人用户端
	if customerSession, exist := model.GetSessionByName(r.CurrentCustomerID); exist {
		customerSession.SendMessage(msg)
	}
}
