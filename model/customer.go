package model

import (
	"fmt"
	"log"
	"reflect"

	avro "com.lueey.shop/protocol"
	"com.lueey.shop/utils"
)

var AllCustomerContainer = map[string]*Customer{}

func GenerateCustomerKey(ID string) string {
	cKey := fmt.Sprintf("customer##%s", ID)
	return cKey
}

type Customer struct {
	ID              string                       // 身份证号
	SalesAdvisorID  string                       // 销售id
	CustomerInfo    *avro.MessageCustomersInfo   // 用户的基本信息
	AuctionGoodsIDs []int32                      // 竞拍得到的孤品
	AuctionRecords  []*avro.MessageAuctionRecord // 竞拍的出价记录
}

func (c *Customer) AddAuctionGoods(goodsID int32) {
	c.AuctionGoodsIDs = append(c.AuctionGoodsIDs, goodsID)
	lang, err := json.Marshal(c.AuctionGoodsIDs)
	if err == nil {
		HSetRedis(GenerateCustomerKey(c.ID), "auctionGoodsIDs", string(lang))
		log.Printf("Update customer auction goods list To %s by idcard %s", string(lang), c.ID)
	}
}

func (c *Customer) UpdateAuctionGoods(goodsIDs []int32) {
	c.AuctionGoodsIDs = goodsIDs
	lang, err := json.Marshal(c.AuctionGoodsIDs)
	if err == nil {
		HSetRedis(GenerateCustomerKey(c.ID), "auctionGoodsIDs", string(lang))
		log.Printf("Update customer auction goods list To %s by idcard %s", string(lang), c.ID)
	}
}

func (c *Customer) AddAuctionRecord(r *avro.MessageAuctionRecord) {
	c.AuctionRecords = append(c.AuctionRecords, r)
	lang, err := json.Marshal(c.AuctionRecords)
	if err == nil {
		HSetRedis(GenerateCustomerKey(c.ID), "auctionRecords", string(lang))
		log.Printf("Update customer auction record list To %s by idcard %s", string(lang), c.ID)
	}
}

func (c *Customer) UpdateAuctionRecord(rs []*avro.MessageAuctionRecord) {
	c.AuctionRecords = rs
	lang, err := json.Marshal(c.AuctionRecords)
	if err == nil {
		HSetRedis(GenerateCustomerKey(c.ID), "auctionRecords", string(lang))
		log.Printf("Update customer auction record list To %s by idcard %s", string(lang), c.ID)
	}
}

func (r *Customer) String() string {
	lang, err := json.MarshalIndent(r, "", "   ")
	if err == nil {
		return string(lang)
	}
	return ""
}

func InitCustomer() {
	// TODO load all customer config
	preload := []string{"110200399833338292", "110182199005170013"}
	for _, ID := range preload {
		cKey := GenerateCustomerKey(ID)
		log.Println(cKey)
		if val, err := HGetRedis(cKey, "IDCard"); err == nil {
			log.Println("IDCard", err, reflect.TypeOf(val), val)
			customerInstance := &Customer{
				ID:           val.(string),
				CustomerInfo: avro.NewMessageCustomersInfo(),
			}
			if val, err := HGetRedis(cKey, "SalesAdvisorID"); err == nil { // 销售人员
				customerInstance.SalesAdvisorID = val.(string)
			}
			if val, err := HGetRedis(cKey, "mobile"); err == nil { // 手机号码
				customerInstance.CustomerInfo.Mobile = &avro.MobileUnion{
					String:    val.(string),
					UnionType: avro.MobileUnionTypeEnumString,
				}
			}
			// if val, err := HGetRedis(cKey, "mobile"); err == nil { // 手机号归属地
			// 	customerInstance.customerInfo.MobileRegion = avro.Customer_mobile_regionUnion
			// }
			if val, err := HGetRedis(cKey, "username"); err == nil { // 用户名
				customerInstance.CustomerInfo.Username = &avro.UsernameUnion{
					String:    val.(string),
					UnionType: avro.UsernameUnionTypeEnumString,
				}
			}

			customerInstance.AuctionRecords = []*avro.MessageAuctionRecord{} // 竞拍记录
			if val, err := HGetRedis(cKey, "auctionRecords"); err == nil {
				if err := json.Unmarshal([]byte(val.(string)), &customerInstance.AuctionRecords); err == nil {
					log.Println(customerInstance.AuctionRecords)
				}
			}

			AllCustomerContainer[customerInstance.ID] = customerInstance
		} else {
			// TODO 暂时没有想好是在用户登录时候初始化， 还是在 init 时统一初始化（或在其他脚本初始化，服务器不需执行初始化操作）
			customerInstance := &Customer{
				ID: ID,
				CustomerInfo: &avro.MessageCustomersInfo{
					Idcard: &avro.IdcardUnion{
						String:    ID,
						UnionType: avro.IdcardUnionTypeEnumString,
					},
					Mobile: &avro.MobileUnion{
						String:    "",
						UnionType: avro.MobileUnionTypeEnumString,
					},
					// MobileRegion: avro.NewMessageCustomersInfo().MobileRegion,
					Username: &avro.UsernameUnion{
						String:    "Test",
						UnionType: avro.UsernameUnionTypeEnumString,
					},
				},
			}
			HSetRedis(cKey, "IDCard", ID)
			HSetRedis(cKey, "SalesAdvisorID", "00012a1e")
			HSetRedis(cKey, "mobile", "")
			HSetRedis(cKey, "username", "Test")
			customerInstance.UpdateAuctionGoods([]int32{})
			customerInstance.UpdateAuctionRecord([]*avro.MessageAuctionRecord{
				&avro.MessageAuctionRecord{
					Goods_id: 1001,
					Customer_mobile: &avro.Customer_mobileUnion{
						String:    "13651027214",
						UnionType: avro.Customer_mobileUnionTypeEnumString,
					},
					Customer_mobile_region: &avro.Customer_mobile_regionUnion{
						String:    "北京",
						UnionType: avro.Customer_mobile_regionUnionTypeEnumString,
					},
					Customer_idcard: &avro.Customer_idcardUnion{
						String:    ID,
						UnionType: avro.Customer_idcardUnionTypeEnumString,
					},
					Customer_username: &avro.Customer_usernameUnion{
						String:    "Test",
						UnionType: avro.Customer_usernameUnionTypeEnumString,
					},
					Bid_price: 10.32,
					Is_final:  false,
					Timestamp: utils.NowMilliseconds(),
				},
			})

			AllCustomerContainer[customerInstance.ID] = customerInstance
			log.Printf("Create new customer info with cKey: %s", cKey)
		}
	}

	log.Printf("%v", AllCustomerContainer)

}
