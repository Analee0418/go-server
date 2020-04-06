package model

import (
	"fmt"
	"log"
	"time"

	"com.lueey.shop/config"
	avro "com.lueey.shop/protocol"
	"com.lueey.shop/utils"
)

var AllCustomerContainer = map[string]*Customer{}

func GenerateCustomerKey(ID string) string {
	cKey := fmt.Sprintf("customer##%s", ID)
	return cKey
}

type Customer struct {
	ID              string          // 身份证号
	UserName        string          // 用户昵称
	SalesAdvisorID  string          // 所属销售ID
	Mobile          string          // 用户手机号
	MobileRegion    string          // 用户手机号所属地
	Address         string          // 家庭住址
	AuctionRecords  []AuctionRecord // 竞拍记录
	AuctionGoodsIDs []int32         // 竞拍得到的物品ID列表
	SignedContract  bool            // 已签约
}

func (r *Customer) BuildCustomerMessage() *avro.MessageCustomersInfo {
	if currentCustomer, ok := AllCustomerContainer[r.ID]; ok {
		return &avro.MessageCustomersInfo{
			Mobile: &avro.MobileUnion{
				UnionType: avro.MobileUnionTypeEnumString,
				String:    currentCustomer.Mobile,
			},

			MobileRegion: &avro.MobileRegionUnion{
				UnionType: avro.MobileRegionUnionTypeEnumString,
				String:    currentCustomer.MobileRegion,
			},

			Idcard: &avro.IdcardUnion{
				UnionType: avro.IdcardUnionTypeEnumString,
				String:    currentCustomer.ID,
			},

			Username: &avro.UsernameUnion{
				UnionType: avro.UsernameUnionTypeEnumString,
				String:    currentCustomer.UserName,
			},

			Address: &avro.AddressUnion{
				UnionType: avro.AddressUnionTypeEnumString,
				String:    currentCustomer.Address,
			},
		}
	} else {
		return &avro.MessageCustomersInfo{
			Mobile: &avro.MobileUnion{
				UnionType: avro.MobileUnionTypeEnumNull,
			},

			MobileRegion: &avro.MobileRegionUnion{
				UnionType: avro.MobileRegionUnionTypeEnumNull,
			},

			Idcard: &avro.IdcardUnion{
				UnionType: avro.IdcardUnionTypeEnumNull,
			},

			Username: &avro.UsernameUnion{
				UnionType: avro.UsernameUnionTypeEnumNull,
			},

			Address: &avro.AddressUnion{
				UnionType: avro.AddressUnionTypeEnumNull,
			},
		}
	}
}

func (c *Customer) ConfirmAuctionGoods(goodsID int32) {
	if _, ok := AllAuctionGoodsContainer[goodsID]; ok {
		c.AuctionGoodsIDs = append(c.AuctionGoodsIDs, goodsID)
		lang, err := json.Marshal(c.AuctionGoodsIDs)
		if err == nil {
			utils.HSetRedis(GenerateCustomerKey(c.ID), "auctionGoodsIDs", string(lang))
			log.Printf("Update customer auction goods list To %s by idcard %s", string(lang), c.ID)
		}
	}
}

func (c *Customer) BiddingGoods(r *avro.MessageAuctionRecord) {
	c.AuctionRecords = append(c.AuctionRecords, AuctionRecord{
		BidPrice:       r.Bid_price,
		Timestamp:      utils.NowMilliseconds(),
		CustomerIDcard: c.ID,
		GoodsID:        r.Goods_id,
	})

	lang, err := json.Marshal(c.AuctionRecords)
	if err == nil {
		utils.HSetRedis(GenerateCustomerKey(c.ID), "recordList", string(lang))
		log.Printf("Add new user bid info %s to cutomer ID %s", string(lang), c.ID)
	}
}

func (c *Customer) ConfirmedSignContract() {
	c.SignedContract = true
	utils.HSetRedis(GenerateCustomerKey(c.ID), "signedContract", "1")
}

func (r *Customer) String() string {
	lang, err := json.MarshalIndent(r, "", "   ")
	if err == nil {
		return string(lang)
	}
	return ""
}

func InitCustomer() {
	log.Println("Start to load customer data.")
	for ID, template := range config.CustomerTemplate {
		customerInstance := &Customer{
			ID:             ID,
			UserName:       template["username"],
			SalesAdvisorID: template["sales_advisor"],
			Mobile:         template["mobile"],
			MobileRegion:   template["mobile_region"],
			Address:        template["address"],
		}

		cKey := GenerateCustomerKey(ID)
		if val, err := utils.HGetRedis(cKey, "IDCard"); err == nil {
			customerInstance.ID = val.(string)

			if val, err := utils.HGetRedis(cKey, "auctionGoodsIDs"); err == nil {
				if err := json.Unmarshal([]byte(val.(string)), &customerInstance.AuctionGoodsIDs); err == nil {
					log.Println(customerInstance.AuctionGoodsIDs)
				}
			}

			if val, err := utils.HGetRedis(cKey, "recordList"); err == nil {
				if err := json.Unmarshal([]byte(val.(string)), &customerInstance.AuctionRecords); err == nil {
					log.Println(customerInstance.AuctionRecords)
				}
			}

			if val, err := utils.HGetRedis(cKey, "signedContract"); err == nil {
				if val.(string) == "1" {
					customerInstance.SignedContract = true
				}
			}

			AllCustomerContainer[customerInstance.ID] = customerInstance
		} else {
			utils.HSetRedis(cKey, "IDCard", ID)
			utils.HSetRedis(cKey, "username", customerInstance.UserName)
			utils.HSetRedis(cKey, "mobile", customerInstance.Mobile)

			AllCustomerContainer[customerInstance.ID] = customerInstance
			log.Printf("Create new customer info %v with cKey: %s", customerInstance, cKey)
		}
	}

	log.Printf("%v", AllCustomerContainer)
	log.Println("Load customer data OK.")
	time.Sleep(time.Second * 1)

}

func PostInitCustoemr() {
}
