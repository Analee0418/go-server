package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"

	"com.lueey.shop/model"
	avro "com.lueey.shop/protocol"
	"com.lueey.shop/utils"
)

// SalesAdvisorBuildContract 销售端生成合约文件转发给用户端
type SalesAdvisorBuildContract struct {
	HandlerSelector
}

func (h *SalesAdvisorBuildContract) setConn(conn *net.Conn) {
	h.conn = conn
}

func (h *SalesAdvisorBuildContract) selected(s *model.Session) {
	h.session = s
}

func (h *SalesAdvisorBuildContract) do(msg avro.Message) {
	if h.session.Room() == nil {
		log.Println("ERROR: sales roomInfo is nil, please signin first.")
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

	// 通知销售端
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

	//
	ss := strings.Split(msg.Sales_advisor_build_contract.RequestSalesAdvisorBuildContract.Filename, ".")
	err := ioutil.WriteFile(utils.ExpandUser(fmt.Sprintf("~/contract.tmp/conract.%s.%s",
		r.CurrentCustomerID, ss[len(ss)-1])),
		msg.Sales_advisor_build_contract.RequestSalesAdvisorBuildContract.Filebytes, 0644)
	if err != nil {
		log.Printf("ERROR: %v", err)
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    "上传失败请稍后重试",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		h.session.SendMessage(msg)
		return
	}

	customerSession.SendMessage(msg)
}
