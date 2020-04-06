package model

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"sync"

	jsoniter "github.com/json-iterator/go"

	"com.lueey.shop/config"
	avro "com.lueey.shop/protocol"
	"com.lueey.shop/utils"
	guuid "github.com/google/uuid"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var rwm sync.RWMutex
var RoomContainer = map[string]*Room{}

func GenerateRoomKey(advisorID string) string {
	roomKey := fmt.Sprintf("room##%s", advisorID)
	return roomKey
}

type CarModel struct {
	Brand    string // 品牌
	Color    string // 颜色
	Interior string // 内饰
	Series   string // 系列
}

type Room struct {
	UUID              int32
	SalesAdvisorID    string
	OrderCount        int32
	CurrentCustomerID string    // 当前正在交谈的客户
	WaitingList       []string  // 排队信息
	CarModel          *CarModel // 汽车模型信息
}

func (r *Room) BuildRoomMessage() *avro.MessageRoomInfo {
	msg := &avro.MessageRoomInfo{
		Room_id:     r.UUID,
		Order_count: r.OrderCount,
		Customer_auction_info: &avro.MessageCustomersAuctionInfo{
			Auction_list:  avro.NewMapDouble(),
			Discount_list: avro.NewMapDouble(),
		},
	}
	if currentCustomer, ok := AllCustomerContainer[r.CurrentCustomerID]; ok {
		if customrMsg := currentCustomer.BuildCustomerMessage(); customrMsg != nil {
			log.Println("DEBUG: customrMsg: ", customrMsg)
			msg.Customer_info = customrMsg
		}
		for _, gid := range currentCustomer.AuctionGoodsIDs {
			if g, ok := AllAuctionGoodsContainer[gid]; ok {
				msg.Customer_auction_info.Auction_list.M[g.GoodsName] = float64(g.FinalPrice)
			}
		}
	} else {
		null := &Customer{}
		msg.Customer_info = null.BuildCustomerMessage()
	}

	for _, cid := range r.WaitingList {
		if waitingCustomer, ok := AllCustomerContainer[cid]; ok {
			if customrMsg := waitingCustomer.BuildCustomerMessage(); customrMsg != nil {
				lang, err := json.MarshalIndent(customrMsg, "", "   ")
				if err == nil {
					log.Println(string(lang))
				}
				msg.Waiting_list = append(msg.Waiting_list, customrMsg)
			}
		}
	}
	if r.CarModel != nil {
		msg.Car_model = &avro.MessageCarsModel{
			Brand: &avro.BrandUnion{
				UnionType: avro.BrandUnionTypeEnumString,
				String:    r.CarModel.Brand,
			},
			Color: &avro.ColorUnion{
				UnionType: avro.ColorUnionTypeEnumString,
				String:    r.CarModel.Color,
			},
			Interior: &avro.InteriorUnion{
				UnionType: avro.InteriorUnionTypeEnumString,
				String:    r.CarModel.Interior,
			},
			Series: &avro.SeriesUnion{
				UnionType: avro.SeriesUnionTypeEnumString,
				String:    r.CarModel.Series,
			},
		}
	} else {
		msg.Car_model = &avro.MessageCarsModel{
			Brand: &avro.BrandUnion{
				UnionType: avro.BrandUnionTypeEnumNull,
			},
			Color: &avro.ColorUnion{
				UnionType: avro.ColorUnionTypeEnumNull,
			},
			Interior: &avro.InteriorUnion{
				UnionType: avro.InteriorUnionTypeEnumNull,
			},
			Series: &avro.SeriesUnion{
				UnionType: avro.SeriesUnionTypeEnumNull,
			},
		}
	}
	return msg
}

func (r *Room) UpdateRoomID() {
	r.UUID = int32(guuid.New().ID())
	utils.HSetRedis(GenerateRoomKey(r.SalesAdvisorID), "roomID", r.UUID)
	log.Printf("Rebuild roomID by roomkey %s", r.SalesAdvisorID)
}

func (r *Room) UpdateOrderCount(orderCount int32) {
	r.OrderCount = orderCount
	utils.HSetRedis(GenerateRoomKey(r.SalesAdvisorID), "orderCount", r.OrderCount)
	log.Printf("Update order count To %d by roomkey %v", orderCount, r.SalesAdvisorID)
}

func (r *Room) UpdateCustomer(idcard string) {
	r.CurrentCustomerID = idcard
	utils.HSetRedis(GenerateRoomKey(r.SalesAdvisorID), "currentCustomerID", r.CurrentCustomerID)
	log.Printf("Update customer_info To %v by roomkey %v", r.CurrentCustomerID, r.SalesAdvisorID)
}

func (r *Room) NewCustomerJoinWaitingList(customerID string) {
	rwm.RLock()
	defer rwm.RUnlock()
	r.WaitingList = append(r.WaitingList, customerID)
	lang, err := json.Marshal(r.WaitingList)
	if err == nil {
		utils.HSetRedis(GenerateRoomKey(r.SalesAdvisorID), "customerWaitingList", string(lang))
		log.Printf("New member join waiting_list %v by roomkey %v", string(lang), r.SalesAdvisorID)
	}
}

func (r *Room) DeleteWaiting(customerID string) {
	rwm.RLock()
	defer rwm.RUnlock()
	newWaitingList := []string{}
	for _, cid := range r.WaitingList {
		if cid == customerID {
			continue
		}
		newWaitingList = append(newWaitingList, cid)
	}
	r.WaitingList = newWaitingList

	lang, err := json.Marshal(r.WaitingList)
	if err == nil {
		utils.HSetRedis(GenerateRoomKey(r.SalesAdvisorID), "customerWaitingList", string(lang))
		log.Printf("Move member to current from waiting_list %v by roomkey %v", string(lang), r.SalesAdvisorID)
	}
}

func (r *Room) UpdateCarModel(model *avro.MessageCarsModel) {
	r.CarModel.Brand = model.Brand.String
	r.CarModel.Color = model.Color.String
	r.CarModel.Series = model.Series.String
	r.CarModel.Interior = model.Interior.String

	lang, err := json.Marshal(r.CarModel)
	if err == nil {
		utils.HSetRedis(GenerateRoomKey(r.SalesAdvisorID), "carModel", string(lang))
		log.Printf("Update carModel To %v by roomkey %v", string(lang), r.SalesAdvisorID)
	}
}

func (r *Room) String() string {
	// obj := map[string]interface{}{}
	// obj["key"] = r.roomKey
	// obj["ID"] = r.roomInfo.Room_id
	// obj["orderCount"] = r.roomInfo.Order_count

	// fooList := []map[string]interface{}{}
	// for _, foo := range r.roomInfo.Waiting_list {
	// 	fooList = append(fooList, map[string]interface{}{
	// 		"customerMobile": foo.Mobile.String,
	// 		"customerIDCard": foo.Idcard.String,
	// 		"customerName":   foo.Username.String,
	// 	})
	// }
	// lang, err := json.Marshal(fooList)
	// if err == nil {
	// 	obj["waitingList"] = string(lang)
	// }
	// lang, err = json.Marshal(r.roomInfo.Customer_auction_info.Auction_list.M)
	// if err == nil {
	// 	obj["customerAuction"] = string(lang)
	// }
	// lang, err = json.Marshal(r.roomInfo.Customer_auction_info.Discount_list.M)
	// if err == nil {
	// 	obj["customerDiscount"] = string(lang)
	// }

	// obj["carBrand"] = r.roomInfo.Car_model.Brand.String
	// obj["carColor"] = r.roomInfo.Car_model.Color.String
	// obj["carSeries"] = r.roomInfo.Car_model.Color.String

	lang, err := json.MarshalIndent(r, "", "   ")
	if err == nil {
		return string(lang)
	}
	return ""
}

func InitRoom() {
	for advisorID, _ := range config.SalesAdvisorTemplate {
		roomInstance := &Room{
			SalesAdvisorID: advisorID,
		}
		roomKey := GenerateRoomKey(advisorID)
		if val, err := utils.HGetRedis(roomKey, "roomID"); err == nil {
			valInt, err := strconv.ParseInt(val.(string), 10, 32)
			log.Println("roomID", err, reflect.TypeOf(val), val)
			roomInstance.UUID = int32(valInt)

			if val, err := utils.HGetRedis(roomKey, "orderCount"); err == nil {
				valInt, err := strconv.ParseInt(val.(string), 10, 32)
				if err == nil {
					roomInstance.OrderCount = int32(valInt)
				}
			}

			if val, err := utils.HGetRedis(roomKey, "currentCustomerID"); err == nil {
				if err == nil {
					roomInstance.CurrentCustomerID = val.(string)
				}
			}

			if val, err := utils.HGetRedis(roomKey, "customerWaitingList"); err == nil {
				if err := json.Unmarshal([]byte(val.(string)), &roomInstance.WaitingList); err == nil {
					log.Println(roomInstance.WaitingList)
				}
			}

			if val, err := utils.HGetRedis(roomKey, "carModel"); err == nil {
				if err := json.Unmarshal([]byte(val.(string)), &roomInstance.CarModel); err == nil {
					log.Println(roomInstance.CarModel)
				}
			}

			RoomContainer[roomInstance.SalesAdvisorID] = roomInstance
		} else {
			roomInstance.UpdateRoomID()
			RoomContainer[roomInstance.SalesAdvisorID] = roomInstance
			log.Printf("Create new room info with roomKey: %s", roomKey)
		}
	}

	log.Printf("%v", RoomContainer)

}

func PostInitRoom() {

}
