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

// CurrentAuctionGoodsID 当前的正在竞拍的商品
var CurrentAuctionGoodsID int32 = -1

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

	log.Printf("Start to load global data OK.\n\n")
}

func PostInitGlobal() {
	// 已签约的用户
	for id, c := range AllCustomerContainer {
		if c.SignedContract {
			GlobalSignedContract[id] = c
		}
	}

	// 在房间内的用户
	for _, r := range RoomContainer {
		if c, ok := AllCustomerContainer[r.CurrentCustomerID]; ok {
			GlobalInRooms[c.ID] = c
		}

		for _, cid := range r.WaitingList {
			if c, ok := AllCustomerContainer[cid]; ok {
				GlobalInRooms[c.ID] = c
			}
		}
	}

}

func GlobalOnHostsSwitchState(s avro.GlobalState) {
	GlobalState = s
	utils.HSetRedis("global", "state", GlobalState.String())

	// TODO notify all session
}

func GlobalOnHostsChoiceGoods(gid int32) {
	// if goods, ok := AllAuctionGoodsContainer[gid]; ok {
	// 	if goods.FinalRecord == nil
	// 	CurrentAuctionGoodsID = gid
	// }
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
