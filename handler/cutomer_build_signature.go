package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"

	"com.lueey.shop/config"
	"com.lueey.shop/model"
	avro "com.lueey.shop/protocol"
	"com.lueey.shop/utils"
)

// CustomerBuildSignature 用户端生成签名转发给销售端
type CustomerBuildSignature struct {
	HandlerSelector
}

func (h *CustomerBuildSignature) setConn(conn *net.Conn) {
	h.conn = conn
}

func (h *CustomerBuildSignature) selected(s *model.Session) {
	h.session = s
}

func (h *CustomerBuildSignature) do(msg avro.Message) {
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
			String:    "您已成功签约，无需再次操作",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		h.session.SendMessage(msg)
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

	// 通知销售端
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

	//
	ss := strings.Split(msg.Request_customer_build_signature.RequestCustomerBuildSignature.Filename, ".")
	err := ioutil.WriteFile(utils.ExpandUser(fmt.Sprintf("~/contract.tmp/signature.%s.%s",
		h.session.CurrentUser().ID, ss[len(ss)-1])),
		msg.Request_customer_build_signature.RequestCustomerBuildSignature.Filebytes, 0644)
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

	salesAdvisorSession.SendMessage(msg)

}
