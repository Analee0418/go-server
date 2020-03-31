package model

import (
	"fmt"
	"log"
	"reflect"
	"strconv"

	jsoniter "github.com/json-iterator/go"

	avro "com.lueey.shop/protocol"
	guuid "github.com/google/uuid"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var AdvisorID = []string{"00012a1e", "c00000e880"}
var RoomContainer = map[string]*Room{}

func GenerateRoomKey(advisorID string) string {
	roomKey := fmt.Sprintf("room##%s", advisorID)
	return roomKey
}

type Room struct {
	roomKey  string
	roomInfo *avro.MessageRoomInfo
}

func (r *Room) GetRoomInfo() avro.MessageRoomInfo {
	return *r.roomInfo
}

func (r *Room) UpdateRoomID() {
	r.roomInfo.Room_id = int32(guuid.New().ID())
	HSetRedis(r.roomKey, "roomID", r.roomInfo.Room_id)
	log.Printf("Rebuild roomID by roomkey %v", r.roomKey)
}

func (r *Room) UpdateOrderCount(orderCount int32) {
	r.roomInfo.Order_count = orderCount
	HSetRedis(r.roomKey, "orderCount", r.roomInfo.Order_count)
	log.Printf("Update order count To %d by roomkey %v", orderCount, r.roomKey)
}

func (r *Room) UpdateCustomerInfo(customer *avro.MessageCustomersInfo) {
	r.roomInfo.Customer_info = customer
	HSetRedis(r.roomKey, "customerMobile", customer.Mobile.String)
	HSetRedis(r.roomKey, "customerIDCard", customer.Idcard.String)
	HSetRedis(r.roomKey, "customerName", customer.Username.String)

	log.Printf("Update customer_info To %v by roomkey %v",
		[3]string{customer.Mobile.String, customer.Idcard.String, customer.Username.String},
		r.roomKey)
}

func (r *Room) UpdateWaitingList(_list []*avro.MessageCustomersInfo) {
	r.roomInfo.Waiting_list = _list
	fooList := []map[string]interface{}{}
	for _, foo := range r.roomInfo.Waiting_list {
		fooList = append(fooList, map[string]interface{}{
			"customerMobile": foo.Mobile.String,
			"customerIDCard": foo.Idcard.String,
			"customerName":   foo.Username.String,
		})
	}
	lang, err := json.Marshal(fooList)
	if err == nil {
		HSetRedis(r.roomKey, "customerWaitingList", string(lang))
		log.Printf("Update waiting_list To %v by roomkey %v", string(lang), r.roomKey)
	}
}

func (r *Room) UpdateCustomerAuction(auction *avro.MessageCustomersAuctionInfo) {
	r.roomInfo.Customer_auction_info = auction
	lang, err := json.Marshal(r.roomInfo.Customer_auction_info.Auction_list.M)
	if err == nil {
		HSetRedis(r.roomKey, "customerAuction", string(lang))
		log.Printf("Update customerAuction To %s by roomkey %v", string(lang), r.roomKey)
	}
	lang, err = json.Marshal(r.roomInfo.Customer_auction_info.Discount_list.M)
	if err == nil {
		HSetRedis(r.roomKey, "customerDiscount", string(lang))
		log.Printf("Update customerDiscount To %s by roomkey %v", string(lang), r.roomKey)
	}
}

func (r *Room) UpdateCarModel(model *avro.MessageCarsModel) {
	r.roomInfo.Car_model = model
	HSetRedis(r.roomKey, "carBrand", r.roomInfo.Car_model.Brand.String)
	HSetRedis(r.roomKey, "carColor", r.roomInfo.Car_model.Color.String)
	HSetRedis(r.roomKey, "carSeries", r.roomInfo.Car_model.Series.String)

	log.Printf("Update car model To %v by roomkey %v",
		[3]string{
			r.roomInfo.Car_model.Brand.String,
			r.roomInfo.Car_model.Color.String,
			r.roomInfo.Car_model.Series.String},
		r.roomKey)
}

func (r *Room) String() string {
	obj := map[string]interface{}{}
	obj["key"] = r.roomKey
	obj["ID"] = r.roomInfo.Room_id
	obj["orderCount"] = r.roomInfo.Order_count

	fooList := []map[string]interface{}{}
	for _, foo := range r.roomInfo.Waiting_list {
		fooList = append(fooList, map[string]interface{}{
			"customerMobile": foo.Mobile.String,
			"customerIDCard": foo.Idcard.String,
			"customerName":   foo.Username.String,
		})
	}
	lang, err := json.Marshal(fooList)
	if err == nil {
		obj["waitingList"] = string(lang)
	}
	lang, err = json.Marshal(r.roomInfo.Customer_auction_info.Auction_list.M)
	if err == nil {
		obj["customerAuction"] = string(lang)
	}
	lang, err = json.Marshal(r.roomInfo.Customer_auction_info.Discount_list.M)
	if err == nil {
		obj["customerDiscount"] = string(lang)
	}

	obj["carBrand"] = r.roomInfo.Car_model.Brand.String
	obj["carColor"] = r.roomInfo.Car_model.Color.String
	obj["carSeries"] = r.roomInfo.Car_model.Color.String

	lang, err = json.MarshalIndent(obj, "", "   ")
	if err == nil {
		strs := string(lang)
		sts := fmt.Sprintf("[roomKey: %s, roomID: %v]\n%s", r.roomKey, r.roomInfo.Room_id, strs)
		return sts
	}
	return ""
}

func InitRoom() {
	for _, advisorID := range AdvisorID {
		roomKey := GenerateRoomKey(advisorID)
		log.Println(roomKey)
		if val, err := HGetRedis(roomKey, "roomID"); err == nil {
			valInt, err := strconv.ParseInt(val.(string), 10, 32)
			log.Println("roomID", err, reflect.TypeOf(val), val)
			roomInstance := &Room{roomKey: roomKey, roomInfo: &avro.MessageRoomInfo{
				Room_id: int32(valInt),
			}}
			if val, err := HGetRedis(roomKey, "orderCount"); err == nil { // 成交单数
				valInt, err := strconv.ParseInt(val.(string), 10, 32)
				if err == nil {
					roomInstance.roomInfo.Order_count = int32(valInt)
				}
			}

			roomInstance.roomInfo.Customer_info = &avro.MessageCustomersInfo{
				Mobile:       avro.NewMobileUnion(),
				MobileRegion: avro.NewMobileRegionUnion(),
				Idcard:       avro.NewIdcardUnion(),
				Username:     avro.NewUsernameUnion(),
			} // 客户信息
			if val, err := HGetRedis(roomKey, "customerMobile"); err == nil {
				roomInstance.roomInfo.Customer_info.Mobile.String = val.(string)
			}
			if val, err := HGetRedis(roomKey, "customerIDCard"); err == nil {
				roomInstance.roomInfo.Customer_info.Idcard.String = val.(string)
			}
			if val, err := HGetRedis(roomKey, "customerName"); err == nil {
				roomInstance.roomInfo.Customer_info.Username.String = val.(string)
			}

			roomInstance.roomInfo.Waiting_list = []*avro.MessageCustomersInfo{} // 排队人数
			if val, err := HGetRedis(roomKey, "customerWaitingList"); err == nil {

				fooList := []map[string]interface{}{}
				if err := json.Unmarshal([]byte(val.(string)), &fooList); err == nil {
					for _, item := range fooList {
						roomInstance.roomInfo.Waiting_list = append(roomInstance.roomInfo.Waiting_list, &avro.MessageCustomersInfo{
							Mobile:       &avro.MobileUnion{String: item["customerMobile"].(string), UnionType: avro.MobileUnionTypeEnumString},
							MobileRegion: &avro.MobileRegionUnion{String: "新疆 伊犁哈萨克自治州", UnionType: avro.MobileRegionUnionTypeEnumString},
							Idcard:       &avro.IdcardUnion{String: item["customerIDCard"].(string), UnionType: avro.IdcardUnionTypeEnumString},
							Username:     &avro.UsernameUnion{String: item["customerName"].(string), UnionType: avro.UsernameUnionTypeEnumString},
						})
					}
				}
			}

			roomInstance.roomInfo.Customer_auction_info = &avro.MessageCustomersAuctionInfo{
				Auction_list:  avro.NewMapDouble(),
				Discount_list: avro.NewMapDouble(),
			} // 竞拍信息
			if val, err := HGetRedis(roomKey, "customerAuction"); err == nil {
				fooMap := map[string]float64{}
				if err := json.Unmarshal([]byte(val.(string)), &fooMap); err == nil {
					roomInstance.roomInfo.Customer_auction_info.Auction_list.M = fooMap
				}
			}
			if val, err := HGetRedis(roomKey, "customerDiscount"); err == nil {
				fooMap := map[string]float64{}
				if err := json.Unmarshal([]byte(val.(string)), &fooMap); err == nil {
					roomInstance.roomInfo.Customer_auction_info.Discount_list.M = fooMap
				}
			}

			roomInstance.roomInfo.Car_model = &avro.MessageCarsModel{
				Brand:  avro.NewBrandUnion(),
				Color:  avro.NewColorUnion(),
				Series: avro.NewSeriesUnion(),
			} // 汽车模型
			if val, err := HGetRedis(roomKey, "carBrand"); err == nil {
				roomInstance.roomInfo.Car_model.Brand.String = val.(string)
			}
			if val, err := HGetRedis(roomKey, "carColor"); err == nil {
				roomInstance.roomInfo.Car_model.Color.String = val.(string)
			}
			if val, err := HGetRedis(roomKey, "carSeries"); err == nil {
				roomInstance.roomInfo.Car_model.Series.String = val.(string)
			}

			RoomContainer[roomInstance.roomKey] = roomInstance
		} else {
			// TODO 根据配置文件初始化（或在其他脚本初始化，服务器不需执行初始化操作）
			roomInstance := &Room{roomKey: roomKey, roomInfo: avro.NewMessageRoomInfo()}
			roomInstance.UpdateRoomID()
			roomInstance.UpdateOrderCount(0)                            // 成交数量
			roomInstance.UpdateCustomerInfo(&avro.MessageCustomersInfo{ // 客户信息
				Mobile:       &avro.MobileUnion{String: "13119999999", UnionType: avro.MobileUnionTypeEnumString},
				MobileRegion: &avro.MobileRegionUnion{String: "新疆 伊犁哈萨克自治州", UnionType: avro.MobileRegionUnionTypeEnumString},
				Idcard:       &avro.IdcardUnion{String: "110200399833338292", UnionType: avro.IdcardUnionTypeEnumString},
				Username:     &avro.UsernameUnion{String: "测试用户1", UnionType: avro.UsernameUnionTypeEnumString},
			})
			roomInstance.UpdateWaitingList([]*avro.MessageCustomersInfo{ // 排队人数
				&avro.MessageCustomersInfo{
					Mobile:       &avro.MobileUnion{String: "13119999999", UnionType: avro.MobileUnionTypeEnumString},
					MobileRegion: &avro.MobileRegionUnion{String: "新疆 伊犁哈萨克自治州", UnionType: avro.MobileRegionUnionTypeEnumString},
					Idcard:       &avro.IdcardUnion{String: "110200399833338292", UnionType: avro.IdcardUnionTypeEnumString},
					Username:     &avro.UsernameUnion{String: "测试用户1", UnionType: avro.UsernameUnionTypeEnumString},
				},
				&avro.MessageCustomersInfo{
					Mobile:       &avro.MobileUnion{String: "13188379982", UnionType: avro.MobileUnionTypeEnumString},
					MobileRegion: &avro.MobileRegionUnion{String: "辽宁 丹东市", UnionType: avro.MobileRegionUnionTypeEnumString},
					Idcard:       &avro.IdcardUnion{String: "110182199005170013", UnionType: avro.IdcardUnionTypeEnumString},
					Username:     &avro.UsernameUnion{String: "测试用户2", UnionType: avro.UsernameUnionTypeEnumString},
				},
			})
			roomInstance.UpdateCustomerAuction(&avro.MessageCustomersAuctionInfo{ // 竞拍信息
				Auction_list:  &avro.MapDouble{M: map[string]float64{"竞拍1": 31.03, "竞拍2": 323.42, "竞拍3": 109982.00}},
				Discount_list: &avro.MapDouble{M: map[string]float64{"折扣券1": -1000, "折扣券2": -10}},
			})
			roomInstance.UpdateCarModel(&avro.MessageCarsModel{ // 汽车模型
				Brand:  &avro.BrandUnion{String: "一汽大众（应该使用配置文件）", UnionType: avro.BrandUnionTypeEnumString},
				Color:  &avro.ColorUnion{String: "宝石蓝（应该使用配置文件）", UnionType: avro.ColorUnionTypeEnumNull},
				Series: &avro.SeriesUnion{String: "CC（应该使用配置文件）", UnionType: avro.SeriesUnionTypeEnumString},
			})

			RoomContainer[roomInstance.roomKey] = roomInstance
			log.Printf("Create new room info with roomKey: %s", roomKey)
		}
	}

	log.Printf("%v", RoomContainer)

}
