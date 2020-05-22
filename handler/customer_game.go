package handler

import (
	"log"
	"math/rand"
	"net"
	"strings"

	"com.lueey.shop/config"
	"com.lueey.shop/model"
	avro "com.lueey.shop/protocol"
)

// CustomerStartGameHanlder 开始游戏
type CustomerStartGameHanlder struct {
	HandlerSelector
}

func (h *CustomerStartGameHanlder) setConn(conn *net.Conn) {
	h.conn = conn
}

func (h *CustomerStartGameHanlder) selected(s *model.Session) {
	h.session = s
}

func (h *CustomerStartGameHanlder) do(msg avro.Message) {
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
			String:    "无法开始游戏，您没有受销售顾问的邀请",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		h.session.SendMessage(msg)
		return
	}

	if h.session.CurrentUser().CurrentGameID != "" {
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    "您已在游戏中，请先完成当前游戏",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		h.session.SendMessage(msg)
		return
	}

	// TODO 默认接金币游戏
	config := []string{}
	for i := 0; i < 60/0.3; i++ {
		config = append(config, string(rand.Intn(3)))
	}

	configStr := strings.Join(config, ",")

	gameID := msg.Customer_start_game.RequestCustomerStartGame.GameID

	// 返回金币游戏配置信息
	msg = *model.GenerateMessage(avro.ActionMessage_game_config)
	msg.Message_game_config = &avro.Message_game_configUnion{
		UnionType: avro.Message_game_configUnionTypeEnumMessageGameConfig,
		MessageGameConfig: &avro.MessageGameConfig{
			GameID: &avro.GameIDUnion{
				UnionType: avro.GameIDUnionTypeEnumString,
				String:    gameID,
			},
			Config: &avro.ConfigUnion{
				UnionType: avro.ConfigUnionTypeEnumString,
				String:    configStr,
			},
		},
	}
	h.session.SendMessage(msg)

	// 开始游戏
	h.session.CurrentUser().StartGame(gameID, configStr)
	// 刷新用户信息到前端
	msg = *model.GenerateMessage(avro.ActionMessage_customers_info)
	msg.Message_customer_info = &avro.Message_customer_infoUnion{
		UnionType:            avro.Message_customer_infoUnionTypeEnumMessageCustomersInfo,
		MessageCustomersInfo: h.session.CurrentUser().BuildCustomerMessage(),
	}
	h.session.SendMessage(msg)
}

// CustomerUploadGameScoreHanlder 上传游戏结果
type CustomerUploadGameScoreHanlder struct {
	HandlerSelector
}

func (h *CustomerUploadGameScoreHanlder) setConn(conn *net.Conn) {
	h.conn = conn
}

func (h *CustomerUploadGameScoreHanlder) selected(s *model.Session) {
	h.session = s
}

func (h *CustomerUploadGameScoreHanlder) do(msg avro.Message) {
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
			String:    "无法上传游戏结果，您没有受销售顾问的邀请",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		h.session.SendMessage(msg)
		return
	}

	gameID := msg.Customer_upload_game_score.RequestCustomerUploadGameScore.GameID

	if h.session.CurrentUser().CurrentGameID != gameID || h.session.CurrentUser().CurrentGameID == "" {
		msg := *model.GenerateMessage(avro.ActionError_message)
		msg.Error_message = &avro.Error_messageUnion{
			String:    "游戏ID错误，无法找到您当前正在进行的游戏内容",
			UnionType: avro.Error_messageUnionTypeEnumString,
		}
		h.session.SendMessage(msg)
		return
	}

	// TODO 需要根据实际需求计算积分
	score := 100

	// 改变用户状态
	h.session.CurrentUser().ChangeState(avro.CustomerStateIdle)
	// 保存游戏积分
	h.session.CurrentUser().UploadGameScore(h.session.CurrentUser().CurrentGameID, int32(score))
	// 返回金币游戏配置信息
	msg = *model.GenerateMessage(avro.ActionMessage_game_result)
	msg.Message_game_result = &avro.Message_game_resultUnion{
		UnionType: avro.Message_game_resultUnionTypeEnumMessageGameResult,
		MessageGameResult: &avro.MessageGameResult{
			GameID: &avro.GameIDUnion{
				UnionType: avro.GameIDUnionTypeEnumString,
				String:    gameID,
			},
			Score: int32(score),
		},
	}
	h.session.SendMessage(msg)

	// 刷新用户信息到前端
	msg = *model.GenerateMessage(avro.ActionMessage_customers_info)
	msg.Message_customer_info = &avro.Message_customer_infoUnion{
		UnionType:            avro.Message_customer_infoUnionTypeEnumMessageCustomersInfo,
		MessageCustomersInfo: h.session.CurrentUser().BuildCustomerMessage(),
	}
	h.session.SendMessage(msg)
}
