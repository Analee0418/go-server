package model

import (
	"log"

	"com.lueey.shop/config"
	avro "com.lueey.shop/protocol"
	"com.lueey.shop/utils"
	"github.com/go-redis/redis"
)

// GlobalState 全局状态
var GlobalState avro.GlobalState = avro.GlobalStateAwating_starting

// GlobalOnlines 在线用户
var GlobalOnlines map[string]string = map[string]string{}

// GlobalInRooms 正在房间的
var GlobalInRooms map[string]string = map[string]string{}

// GlobalSignedContract 已经签约
var GlobalSignedContract map[string]string = map[string]string{}

// GlobalGame 正在玩游戏
var GlobalGame map[string]string = map[string]string{}

// GlobalVisitor 正在参观模型
var GlobalVisitor map[string]string = map[string]string{}

func initGlobalState() {
	log.Println("INFO: Start to load global state.")
	if val, err := utils.HGetRedis("global", "state"); err == nil {
		GlobalState, err = avro.NewGlobalStateValue(val.(string))
		if err == redis.Nil {
			GlobalState = avro.GlobalStateAwating_starting
		} else {
			log.Fatal("\033[1;31mERROR: \033[0mglobal state is illegal.", err)
		}
	}

	log.Printf("INFO: Start to load global state[%s] OK.\n\n", GlobalState)
}

// TCPInitGlobal 初始化全局状态
func TCPInitGlobal() {
	initGlobalState()
}

// TCPGlobalOnCustomerSignin 用户登录
func TCPGlobalOnCustomerSignin(customerID string) []*avro.Message {
	// 根据全局状态构造对应的Message
	msgs := []*avro.Message{}

	currentCustomer, ok := AllCustomerContainer[customerID]
	if !ok {
		return msgs
	}

	switch GlobalState {
	case avro.GlobalStateChat_with_advisor: // 洽谈阶段
		if r, ok := RoomContainer[currentCustomer.SalesAdvisorID]; ok {
			roominfo := r.BuildRoomMessage()
			if roominfo == nil {
				log.Printf("\033[1;31mERROR: \033[0mBuildRoomMessage failed, room[%v] is not exists.", r)
			} else {
				msg := GenerateMessage(avro.ActionMessage_room_info)
				msg.Message_room_info = &avro.Message_room_infoUnion{
					MessageRoomInfo: roominfo,
					UnionType:       avro.Message_room_infoUnionTypeEnumMessageRoomInfo,
				}
				msgs = append(msgs, msg)
			}
		}
	case avro.GlobalStateAution: // 竞拍阶段
		log.Println("INFO: Aucion step")
	}

	msg := GenerateMessage(avro.ActionMessage_customers_info)
	msg.Message_customer_info = &avro.Message_customer_infoUnion{
		MessageCustomersInfo: currentCustomer.BuildCustomerMessage(),
		UnionType:            avro.Message_customer_infoUnionTypeEnumMessageCustomersInfo,
	}
	msgs = append(msgs, msg)

	return msgs
}

// TCPGlobalReceiveGlobalState 主持人端更新全局状态
func TCPGlobalReceiveGlobalState() {
	pubsub := utils.GetRDB().Subscribe("updated_global_state")
	defer func() { pubsub.Close() }()
	ch := pubsub.ChannelSize(1)
	for {
		res := <-ch
		state := res.Payload
		if r, err := avro.NewGlobalStateValue(state); err == nil {
			msg := GenerateMessage(avro.ActionMessage_global_state)
			msg.Message_global_state = &avro.Message_global_stateUnion{
				UnionType: avro.Message_global_stateUnionTypeEnumMessageGlobalState,
				MessageGlobalState: &avro.MessageGlobalState{
					GlobalState: r,
				},
			}
			log.Printf("INFO: global state changed To %s.", r)
			BroadcastMessage(*msg)
		}
	}
}

// HTTPInitGlobal TODO
func HTTPInitGlobal() {
	//
	initGlobalState()

	// 已签约的用户
	for cardID := range config.CustomerTemplate {
		cKey := GenerateCustomerKey(cardID)
		if _, err := utils.HGetRedis(cKey, "IDCard"); err == nil {
			if val, err := utils.HGetRedis(cKey, "signedContract"); err == nil {
				if val.(string) == "1" {
					GlobalSignedContract[cardID] = cardID
				}
			}
		}
	}

	// 在房间内的用户
	for advisorID := range config.SalesAdvisorTemplate {
		roomKey := GenerateRoomKey(advisorID)
		if _, err := utils.HGetRedis(roomKey, "roomID"); err == nil {
			// 正在洽谈的用户
			if val, err := utils.HGetRedis(roomKey, "currentCustomerID"); err == nil {
				if err == nil {
					customerCardID := val.(string)
					if _, ok := config.CustomerTemplate[customerCardID]; ok {
						GlobalInRooms[customerCardID] = customerCardID
					}
				}
			}
			// 正在排队的用户
			if val, err := utils.HGetRedis(roomKey, "customerWaitingList"); err == nil {
				var WaitingList []string
				if err := json.Unmarshal([]byte(val.(string)), &WaitingList); err == nil {
					for _, customerCardID := range WaitingList {
						if _, ok := config.CustomerTemplate[customerCardID]; ok {
							GlobalInRooms[customerCardID] = customerCardID
						}
					}
				}
			}

		}
	}
}
