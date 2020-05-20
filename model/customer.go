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
	ID                string           // 身份证号
	UserName          string           // 用户昵称
	SalesAdvisorID    string           // 所属销售ID
	Mobile            string           // 用户手机号
	MobileRegion      string           // 用户手机号所属地
	Address           string           // 家庭住址
	AuctionRecords    []AuctionRecord  // 竞拍记录
	AuctionGoodsIDs   []int32          // 竞拍得到的物品ID列表
	SignedContract    bool             // 已签约
	Contract          Contract         // 合约信息
	State             string           // 当前状态
	CurrentGameID     string           // 当前游戏ID
	currentGameConfig string           // 当前游戏配置
	GameData          map[string]int32 // 游戏数据
}

func (r *Customer) BuildCustomerMessage() *avro.MessageCustomersInfo {
	if currentCustomer, ok := AllCustomerContainer[r.ID]; ok {
		msg := &avro.MessageCustomersInfo{
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

			CurrentGameID: &avro.CurrentGameIDUnion{
				UnionType: avro.CurrentGameIDUnionTypeEnumString,
				String:    currentCustomer.CurrentGameID,
			},

			CurrentGameConfig: &avro.CurrentGameConfigUnion{
				UnionType: avro.CurrentGameConfigUnionTypeEnumString,
				String:    currentCustomer.currentGameConfig,
			},
		}
		if s, err := avro.NewCustomerStateValue(currentCustomer.State); err == nil {
			msg.State = s
		}
		return msg
	}

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

		CurrentGameID: &avro.CurrentGameIDUnion{
			UnionType: avro.CurrentGameIDUnionTypeEnumNull,
		},

		CurrentGameConfig: &avro.CurrentGameConfigUnion{
			UnionType: avro.CurrentGameConfigUnionTypeEnumNull,
		},

		State: avro.CustomerStateIdle,
	}
}

func (c *Customer) ConfirmAuctionGoods(goodsID int32) {
	if _, ok := AllAuctionGoodsContainer[goodsID]; ok {
		c.AuctionGoodsIDs = append(c.AuctionGoodsIDs, goodsID)
		lang, err := json.Marshal(c.AuctionGoodsIDs)
		if err == nil {
			utils.HSetRedis(GenerateCustomerKey(c.ID), "auctionGoodsIDs", string(lang))
			log.Printf("[INFO] Update customer auction goods list To %s by idcard %s", string(lang), c.ID)
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
		log.Printf("[INFO] Add new user bid info %s to cutomer ID %s", string(lang), c.ID)
	}
}

func (c *Customer) ConfirmedSignContract(salesID string, price float32, disprice float32, brand string, color string, interior string, series string) {
	c.SignedContract = true
	utils.HSetRedis(GenerateCustomerKey(c.ID), "signedContract", "1")
	c.Contract = CreateContract(c.ID, salesID, price, disprice, brand, color, interior, series)
	log.Printf("\033[1;36mSTATS: \033[0mcustomer[%s] completed contract[%s].", c.ID, c.Contract)
	log.Printf("[INFO] customer[%s] completed contract[%s].", c.ID, c.Contract)
}

func (c *Customer) ChangeState(s avro.CustomerState) {
	c.State = string(s)
	utils.HSetRedis(GenerateCustomerKey(c.ID), "state", string(s))
	log.Printf("\033[1;36mSTATS: \033[0mcustomer[%s] update current state To [%s].", c.ID, c.State)
	log.Printf("[INFO] customer[%s] update current state To [%s].", c.ID, c.State)
}

func (c *Customer) StartGame(gameID string, config string) {
	c.CurrentGameID = gameID
	c.currentGameConfig = config
	utils.HSetRedis(GenerateCustomerKey(c.ID), "currentGameID", gameID)
	utils.HSetRedis(GenerateCustomerKey(c.ID), "currentGameConfig", config)
	log.Printf("\033[1;36mSTATS: \033[0mcustomer[%s] start game [%s, %s].", c.ID, c.CurrentGameID, c.currentGameConfig)
	log.Printf("[INFO] customer[%s] start game [%s, %s].", c.ID, c.CurrentGameID, c.currentGameConfig)
}

func (c *Customer) UploadGameScore(gameID string, score int32) {
	c.CurrentGameID = ""
	utils.HSetRedis(GenerateCustomerKey(c.ID), "currentGameID", "")
	c.currentGameConfig = ""
	utils.HSetRedis(GenerateCustomerKey(c.ID), "currentGameConfig", "")
	c.GameData[gameID] = score
	lang, err := json.Marshal(c.GameData)
	if err == nil {
		utils.HSetRedis(GenerateCustomerKey(c.ID), "gameData", string(lang))
		log.Printf("\033[1;36mSTATS: \033[0mcustomer[%s] upload game score [%s:%d].", c.ID, gameID, score)
		log.Printf("[INFO] customer[%s] upload game score [%s:%d].", c.ID, gameID, score)
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
	log.Println("[INFO] Start to load customer data.")
	for ID, template := range config.CustomerTemplate {
		customerInstance := &Customer{
			ID:             ID,
			UserName:       template["username"],
			SalesAdvisorID: template["sales_advisor"],
			Mobile:         template["mobile"],
			MobileRegion:   template["mobile_region"],
			Address:        template["address"],
			State:          "idle",
			GameData:       make(map[string]int32),
		}

		cKey := GenerateCustomerKey(ID)
		if val, err := utils.HGetRedis(cKey, "IDCard"); err == nil {
			customerInstance.ID = val.(string)

			if val, err := utils.HGetRedis(cKey, "auctionGoodsIDs"); err == nil {
				if err := json.Unmarshal([]byte(val.(string)), &customerInstance.AuctionGoodsIDs); err == nil {
					log.Printf("[INFO] %v", customerInstance.AuctionGoodsIDs)
				}
			}

			if val, err := utils.HGetRedis(cKey, "recordList"); err == nil {
				if err := json.Unmarshal([]byte(val.(string)), &customerInstance.AuctionRecords); err == nil {
					log.Printf("[INFO] %v", customerInstance.AuctionRecords)
				}
			}

			if val, err := utils.HGetRedis(cKey, "signedContract"); err == nil {
				if val.(string) == "1" {
					customerInstance.SignedContract = true
				} else {
					customerInstance.SignedContract = false
				}
			}

			if val, err := utils.HGetRedis(cKey, "state"); err == nil {
				customerInstance.State = val.(string)
			}

			if val, err := utils.HGetRedis(cKey, "currentGameID"); err == nil {
				customerInstance.CurrentGameID = val.(string)
			}

			if val, err := utils.HGetRedis(cKey, "currentGameConfig"); err == nil {
				customerInstance.currentGameConfig = val.(string)
			}

			if val, err := utils.HGetRedis(cKey, "gameData"); err == nil {
				if err := json.Unmarshal([]byte(val.(string)), &customerInstance.GameData); err == nil {
					log.Printf("[INFO] %v", customerInstance.GameData)
				}
			}

			if val, err := utils.HGetRedis(cKey, "dataContract"); err == nil {
				if err := json.Unmarshal([]byte(val.(string)), &customerInstance.Contract); err == nil {
					log.Printf("[INFO] load contract: %s", customerInstance.Contract)
				}
			}

			AllCustomerContainer[customerInstance.ID] = customerInstance
		} else {
			utils.HSetRedis(cKey, "IDCard", ID)
			utils.HSetRedis(cKey, "username", customerInstance.UserName)
			utils.HSetRedis(cKey, "mobile", customerInstance.Mobile)

			AllCustomerContainer[customerInstance.ID] = customerInstance
			log.Printf("\033[1;36mSTATS: \033[0mCreate new customer info %v with cKey: %s, To Redis.", customerInstance, cKey)
			log.Printf("[INFO] Create new customer info %v with cKey: %s, To Redis.", customerInstance, cKey)
		}
	}

	if config.DEBUG {
		log.Printf("[DEBUG] All customer entity: %v", AllCustomerContainer)
	}
	log.Printf("[INFO] Load customer data OK.\n\n")
	time.Sleep(time.Second * 1)

}

func PostInitCustoemr() {
}
