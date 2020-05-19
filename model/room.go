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
	Brand    string  // 品牌
	Color    string  // 颜色
	Interior string  // 内饰
	Series   string  // 系列
	Price    float32 // 价格
}

type Room struct {
	UUID               int32
	SalesAdvisorID     string
	SalesAdvisorMobile string
	SalesAdvisorName   string
	OrderCount         int32
	CurrentCustomerID  string    // 当前正在交谈的客户
	WaitingList        []string  // 排队信息
	CarModel           *CarModel // 汽车模型信息
	Province           string    // 省
	City               string    // 市
	Company            string    // 经销商
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
			if config.DEBUG {
				log.Println("DEBUG: customrMsg: ", customrMsg)
			}
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
	log.Printf("INFO: Rebuild roomID by roomkey[%d] for %s", r.UUID, r.SalesAdvisorID)
}

func (r *Room) UpdateOrderCount(orderCount int32) {
	rwm.RLock()
	defer rwm.RUnlock()
	r.OrderCount = orderCount
	utils.HSetRedis(GenerateRoomKey(r.SalesAdvisorID), "orderCount", r.OrderCount)
	log.Printf("\033[1;36mSTATS: \033[0mUpdate order count To %d by roomID: %d, salesAdvisorID: %v", orderCount, r.UUID, r.SalesAdvisorID)
	log.Printf("INFO: Update order count To %d by roomkey %v", orderCount, r.SalesAdvisorID)
}

func (r *Room) UpdateCustomer(idcard string) {
	r.CurrentCustomerID = idcard
	utils.HSetRedis(GenerateRoomKey(r.SalesAdvisorID), "currentCustomerID", r.CurrentCustomerID)
	log.Printf("\033[1;36mSTATS: \033[0mUpdate customer_info To %s by roomkey %v", r.CurrentCustomerID, r.SalesAdvisorID)
	log.Printf("INFO: Update customer_info To %v by roomkey %v", r.CurrentCustomerID, r.SalesAdvisorID)
}

func (r *Room) NewCustomerJoinWaitingList(customerID string) {
	rwm.RLock()
	defer rwm.RUnlock()
	r.WaitingList = append(r.WaitingList, customerID)
	lang, err := json.Marshal(r.WaitingList)
	if err == nil {
		utils.HSetRedis(GenerateRoomKey(r.SalesAdvisorID), "customerWaitingList", string(lang))
		log.Printf("\033[1;36mSTATS: \033[0mNew member %v join waiting_list %v by roomkey %v", customerID, string(lang), r.SalesAdvisorID)
		log.Printf("INFO: New member join waiting_list %v by roomkey %v", string(lang), r.SalesAdvisorID)
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
		log.Printf("INFO: Move member to room from waiting_list %v by roomkey %v", string(lang), r.SalesAdvisorID)
	}
}

func (r *Room) UpdateCarModel(model *avro.MessageCarsModel) {
	if r.CarModel == nil {
		r.CarModel = &CarModel{}
	}
	if model.Brand.UnionType != avro.BrandUnionTypeEnumNull {
		log.Printf("\033[1;36mSTATS: \033[0msalesAdvisor[%s] change carModel.Brand [%s] -> [%s]", r.SalesAdvisorID, r.CarModel.Brand, model.Brand.String)
		r.CarModel.Brand = model.Brand.String
	}
	if model.Color.UnionType != avro.ColorUnionTypeEnumNull {
		log.Printf("\033[1;36mSTATS: \033[0msalesAdvisor[%s] change carModel.Color [%s] -> [%s]", r.SalesAdvisorID, r.CarModel.Color, model.Color.String)
		r.CarModel.Color = model.Color.String
	}
	if model.Series.UnionType != avro.SeriesUnionTypeEnumNull {
		log.Printf("\033[1;36mSTATS: \033[0msalesAdvisor[%s] change carModel.Series [%s] -> [%s]", r.SalesAdvisorID, r.CarModel.Series, model.Series.String)
		r.CarModel.Series = model.Series.String
	}
	if model.Interior.UnionType != avro.InteriorUnionTypeEnumNull {
		log.Printf("\033[1;36mSTATS: \033[0msalesAdvisor[%s] change carModel.Interior [%s] -> [%s]", r.SalesAdvisorID, r.CarModel.Interior, model.Interior.String)
		r.CarModel.Interior = model.Interior.String
	}

	lang, err := json.Marshal(r.CarModel)
	if err == nil {
		utils.HSetRedis(GenerateRoomKey(r.SalesAdvisorID), "carModel", string(lang))
		log.Printf("INFO: Update carModel To %v by roomkey %v", string(lang), r.SalesAdvisorID)
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
	for advisorID, template := range config.SalesAdvisorTemplate {
		roomInstance := &Room{
			SalesAdvisorID:     advisorID,
			SalesAdvisorMobile: template["advisor_mobile"],
			SalesAdvisorName:   template["advisor_name"],
			Province:           template["advisor_province"],
			City:               template["advisor_city"],
			Company:            template["advisor_company"],
		}
		roomKey := GenerateRoomKey(advisorID)
		if val, err := utils.HGetRedis(roomKey, "roomID"); err == nil {
			valInt, err := strconv.ParseInt(val.(string), 10, 32)
			log.Println("INFO: roomID", err, reflect.TypeOf(val), val)
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
					log.Println("INFO: ", roomInstance.WaitingList)
				}
			}

			if val, err := utils.HGetRedis(roomKey, "carModel"); err == nil {
				if err := json.Unmarshal([]byte(val.(string)), &roomInstance.CarModel); err == nil {
					log.Println("INFO: ", roomInstance.CarModel)
				}
			}

			RoomContainer[roomInstance.SalesAdvisorID] = roomInstance
		} else {
			roomInstance.UpdateRoomID()
			RoomContainer[roomInstance.SalesAdvisorID] = roomInstance
			log.Printf("\033[1;36mSTATS: \033[0mCreate new room info with roomKey: %s\n%s\n", roomKey, roomInstance)
			log.Printf("INFO: Create new room info with roomKey: %s", roomKey)
		}
	}

	log.Printf("INFO: All room entity: %v", RoomContainer)
	log.Printf("INFO: Load room entity data OK.\n\n")

}

func PostInitRoom() {

}
