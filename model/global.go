package model

import (
	"log"

	avro "com.lueey.shop/protocol"
	"com.lueey.shop/utils"
	"github.com/go-redis/redis"
)

// GlobalState 全局状态
var GlobalState avro.GlobalState = avro.GlobalStateAwating_starting

// GlobalOnlines 在线用户
var GlobalOnlines map[string]*Customer = map[string]*Customer{}

// GlobalInRooms 正在房间的
var GlobalInRooms map[string]*Customer = map[string]*Customer{}

// GlobalSignedContract 已经签约
var GlobalSignedContract map[string]*Customer = map[string]*Customer{}

// GlobalGame 正在玩游戏
var GlobalGame map[string]*Customer = map[string]*Customer{}

// GlobalVisitor 正在参观模型
var GlobalVisitor map[string]*Customer = map[string]*Customer{}

// InitGlobal 初始化全局状态
func InitGlobal() {
	log.Println("Start to load global data.")
	if val, err := utils.HGetRedis("global", "state"); err == nil {
		GlobalState, err = avro.NewGlobalStateValue(val.(string))
		if err == redis.Nil {
			GlobalState = avro.GlobalStateAwating_starting
		} else {
			log.Fatal("ERROR: global state is illegal.", err)
		}
	}

	// if val, err := utils.HGetRedis("global", "signedContract"); err == nil {
	// 	GlobalState, err = avro.NewGlobalStateValue(val.(string))
	// 	if err == redis.Nil {
	// 		GlobalState = avro.GlobalStateAwating_starting
	// 	} else {
	// 		log.Fatal("ERROR: global state is illegal.", err)
	// 	}
	// }

	log.Println("Start to load global data OK.")
}

func PostInitGlobal() {

}

func GlobalOnHostsSwitchState(s avro.GlobalState) {
	GlobalState = s
	utils.HSetRedis("global", "state", GlobalState.String())

	// TODO notify all session
}

// GlobalOnCustomerSignin 用户登录
func GlobalOnCustomerSignin(customerID string) []*avro.Message {
	currentCustomer, ok := AllCustomerContainer[customerID]
	if ok {
		GlobalOnlines[customerID] = currentCustomer
	}

	// 根据全局状态构造对应的Message
	msgs := []*avro.Message{}

	switch GlobalState {
	case avro.GlobalStateChat_with_advisor: // 洽谈阶段
		if r, ok := RoomContainer[currentCustomer.SalesAdvisorID]; ok {
			roominfo := r.BuildRoomMessage()
			if roominfo == nil {
				log.Printf("ERROR: BuildRoomMessage failed, room: %v", r)
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
		log.Println("Aucion step")
	}

	return msgs
}

// GlobalOnCustomerDisconnect 用户离线
func GlobalOnCustomerDisconnect(customerID string) {
	if _, ok := GlobalOnlines[customerID]; ok {
		delete(GlobalOnlines, customerID)
	}
}
