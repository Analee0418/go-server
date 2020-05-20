package model

import (
	"strconv"
	"time"

	"fmt"
	"log"

	"com.lueey.shop/config"
	avro "com.lueey.shop/protocol"
	"com.lueey.shop/utils"
)

var AllAuctionGoodsContainer = map[int32]*AuctionGoods{}

func GenerateAuctionGoodsKey(ID int32) string {
	cKey := fmt.Sprintf("auctionGoods##%d", ID)
	return cKey
}

type AuctionRecord struct {
	GoodsID        int32
	CustomerIDcard string
	BidPrice       float32
	Timestamp      int64
}

type AuctionGoods struct {
	GoodsID         int32
	GoodsName       string
	OriginalPrice   float32
	LimitPrice      float32
	FinalPrice      float32
	FinalRecord     *AuctionRecord
	Records         []AuctionRecord
	RecordsUser     map[string]*Customer
	CountdownSecond int32
}

func (r *AuctionRecord) BuildAuctionRecord() *avro.MessageAuctionRecord {
	if c, ok := AllCustomerContainer[r.CustomerIDcard]; ok {
		return &avro.MessageAuctionRecord{
			Goods_id:  r.GoodsID,
			Bid_price: r.BidPrice,
			Timestamp: r.Timestamp,
			Customer_mobile: &avro.Customer_mobileUnion{
				UnionType: avro.Customer_mobileUnionTypeEnumString,
				String:    c.Mobile,
			},
			Customer_mobile_region: &avro.Customer_mobile_regionUnion{
				UnionType: avro.Customer_mobile_regionUnionTypeEnumString,
				String:    c.MobileRegion,
			},
			Customer_idcard: &avro.Customer_idcardUnion{
				UnionType: avro.Customer_idcardUnionTypeEnumString,
				String:    c.ID,
			},
			Customer_username: &avro.Customer_usernameUnion{
				UnionType: avro.Customer_usernameUnionTypeEnumString,
				String:    c.UserName,
			},
		}
	}
	return nil
}

func (r *AuctionGoods) BuildAuctionGoodsMessage() *avro.MessageAuctionGoods {
	msg := &avro.MessageAuctionGoods{
		Goods_id: r.GoodsID,
		Goods_name: &avro.Goods_nameUnion{
			UnionType: avro.Goods_nameUnionTypeEnumString,
			String:    r.GoodsName,
		},
		Limit_price:    r.LimitPrice,
		Original_price: r.OriginalPrice,
		Final_record: &avro.Final_recordUnion{
			UnionType: avro.Final_recordUnionTypeEnumNull,
		},
	}
	if g, ok := AllAuctionGoodsContainer[r.GoodsID]; ok {
		msg.Final_price = g.FinalPrice
		msg.Users_num = int32(len(r.RecordsUser))
		if g.FinalRecord != nil {
			msg.Final_record = &avro.Final_recordUnion{
				UnionType:            avro.Final_recordUnionTypeEnumMessageAuctionRecord,
				MessageAuctionRecord: g.FinalRecord.BuildAuctionRecord(),
			}
		}

		for idx, r := range g.Records {
			if idx >= 5 {
				break
			}
			if rMsg := r.BuildAuctionRecord(); rMsg != nil {
				msg.Auction_records = append(msg.Auction_records, rMsg)
			}
		}

		return msg
	}
	return nil
}

func (g *AuctionGoods) ConfirmFinalOnEndOfAuction(r AuctionRecord) {
	// sort.Slice(g.Records, func(p, q int) bool {
	// 	return g.Records[p].Timestamp < g.Records[q].Timestamp
	// })

	// r := g.Records[len(g.Records)-1]

	g.FinalRecord = &r
	g.FinalPrice = r.BidPrice

	utils.HSetRedis(GenerateAuctionGoodsKey(g.GoodsID), "finalPrice", r.BidPrice)
	lang, err := json.Marshal(g.FinalRecord)
	if err == nil {
		utils.HSetRedis(GenerateAuctionGoodsKey(g.GoodsID), "finalRecord", string(lang))
		log.Printf("[INFO] Confirmed final bid info %s by goodsID %d", string(lang), g.GoodsID)
	}
}

func (g *AuctionGoods) CustomerBidding(customerID string, r avro.MessageAuctionRecord) {
	g.Records = append(g.Records, AuctionRecord{
		BidPrice:       r.Bid_price,
		Timestamp:      utils.NowMilliseconds(),
		CustomerIDcard: customerID,
		GoodsID:        r.Goods_id,
	})

	lang, err := json.Marshal(g.Records)
	if err == nil {
		utils.HSetRedis(GenerateAuctionGoodsKey(g.GoodsID), "recordList", string(lang))
		log.Printf("[INFO] Add new user bid info %s to goodsID %d", string(lang), g.GoodsID)
	}

	if c, ok := AllCustomerContainer[customerID]; ok {
		g.RecordsUser[customerID] = c
	}
}

func (g *AuctionGoods) String() string {
	lang, err := json.MarshalIndent(g, "", "   ")
	if err == nil {
		return string(lang)
	}
	return ""
}

func InitAuctionGoods() {
	log.Println("[INFO] Start to load auction goods data.")
	for goodsID, template := range config.AuctionGoodsTemplate {
		goodsInstance := &AuctionGoods{
			GoodsID:         goodsID,
			GoodsName:       template["goods_name"].(string),
			OriginalPrice:   template["original_price"].(float32),
			LimitPrice:      template["limit_price"].(float32),
			CountdownSecond: template["countdown_second"].(int32),
		}

		gKey := GenerateAuctionGoodsKey(goodsID)
		if val, err := utils.HGetRedis(gKey, "goodsID"); err == nil {
			gid, err := strconv.ParseInt(val.(string), 10, 32)
			if err != nil {
				log.Printf("\033[1;31m[ERROR] \033[0mcan not read goodsID %v", val)
				continue
			}
			goodsInstance.GoodsID = int32(gid)

			if val, err := utils.HGetRedis(gKey, "finalPrice"); err == nil { // 最终价格
				finalPrice, err := strconv.ParseFloat(val.(string), 32)
				if err != nil {
					log.Printf("\033[1;31m[ERROR] \033[0mcan not read finalPrice %v by goodsID %d", val, goodsID)
					continue
				}
				goodsInstance.FinalPrice = float32(finalPrice)
			}

			if val, err := utils.HGetRedis(gKey, "finalRecord"); err == nil {
				if err := json.Unmarshal([]byte(val.(string)), &goodsInstance.FinalRecord); err == nil {
					log.Printf("[INFO] %v", goodsInstance.FinalRecord)
				}
			}

			if val, err := utils.HGetRedis(gKey, "recordList"); err == nil {
				if err := json.Unmarshal([]byte(val.(string)), &goodsInstance.Records); err == nil {
					for _, r := range goodsInstance.Records {
						goodsInstance.RecordsUser[r.CustomerIDcard] = nil
					}
				}
			}

			AllAuctionGoodsContainer[goodsInstance.GoodsID] = goodsInstance
		} else {
			utils.HSetRedis(gKey, "goodsID", goodsID)
			AllAuctionGoodsContainer[goodsID] = goodsInstance
			log.Printf("[INFO] Create new auction goods instance with gKey: %s", gKey)
		}
	}

	if config.DEBUG {
		log.Printf("[INFO] All auction goods: %v", AllAuctionGoodsContainer)
	}
	log.Printf("[INFO] Load auction goods data OK.\n\n")
	time.Sleep(time.Second * 1)
}

func PostInitAuctionGoods() {
	for _, g := range AllAuctionGoodsContainer {
		for _, r := range g.Records {
			g.RecordsUser[r.CustomerIDcard] = AllCustomerContainer[r.CustomerIDcard]
		}
	}
}
