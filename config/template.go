package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"com.lueey.shop/utils"
	"github.com/360EntSecGroup-Skylar/excelize"
)

// CustomerTemplate 客户配置
var CustomerTemplate map[string]map[string]string = map[string]map[string]string{}

// SalesAdvisorTemplate 销售人员配置
var SalesAdvisorTemplate map[string]map[string]string = map[string]map[string]string{}

// AuctionGoodsTemplate 竞拍商品配置
var AuctionGoodsTemplate map[int32]map[string]interface{} = map[int32]map[string]interface{}{}

func Init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
	log.Println("Start to load templated.")
	files, _ := filepath.Glob(utils.ExpandUser("~/config/*.xlsx"))
	lang, err := json.MarshalIndent(files, "", "   ")
	if err == nil {
		strs := string(lang)
		log.Printf("==================== ALL CONFIG FILES %s", strs) // contains a list of all files in the current directory

		for _, filename := range files {
			log.Printf("Parse config %v", filename)
			xlsx, err := excelize.OpenFile(filename)
			if err != nil {
				log.Panic(err)
			}

			// Example

			// Get value from cell by given sheet index and axis.
			// cell := xlsx.GetCellValue("Sheet1", "B2")
			// log.Println(cell)

			// Get sheet index.
			// index := xlsx.GetSheetIndex("Sheet1")
			// Get all the rows in a sheet.
			rows := xlsx.GetRows("Sheet1")
			// log.Println(rows[100:])
			if DEBUG {
				log.Println(strings.Join(rows[0], ",\t"))
			}
			lastAdvisor := ""
			lastAdvisorName := ""
			for _, row := range rows[1:] {
				if DEBUG {
					log.Println(strings.Join(row, ",\t"))
				}
				if strings.Contains(filename, "销售") {
					idcard := row[0]
					mobile := row[1]
					region := ""

					if row[5] != "" {
						lastAdvisor = row[5]
					}

					if row[4] != "" {
						lastAdvisorName = row[4]
					}

					if val, err := utils.HGetRedis("mobileRegion", mobile); err == nil {

						region = val.(string)
						log.Printf("Mobile region from redis: %s", region)

					} else {

						resp, err := http.Get("https://www.ip.cn/db?num=%2B" + mobile)
						if err != nil {
							log.Printf("ERROR: crawl mobile region failed. %v", err)
							continue
						}

						if resp.StatusCode == 200 {
							body, err := ioutil.ReadAll(resp.Body)
							if err != nil {
								log.Printf("ERROR: read mobile region failed. %v", err)
								continue
							}
							bodystr := string(body)
							res, _ := regexp.Compile("</code>&nbsp;所在城市: (.*)<br /><br />")
							matched := res.FindAllString(bodystr, -1)
							if matched == nil {
								region = ""
							} else {
								region = strings.Replace(strings.Replace(matched[0], "</code>&nbsp;所在城市: ", "", -1), "<br /><br />", "", -1)
							}

							if region != "" {
								utils.HSetRedis("mobileRegion", mobile, region)
							}
						}
					}

					customerTemplate := map[string]string{
						"idcard":        idcard,
						"mobile":        mobile,
						"mobile_region": region,
						"username":      row[2],
						"address":       row[3],
						"sales_advisor": lastAdvisor,
					}
					CustomerTemplate[idcard] = customerTemplate

					salesadvisorID := row[5]
					salesAdvisorTemplate := map[string]string{
						"advisor_name": lastAdvisorName,
						"advisor_id":   lastAdvisor,
					}
					SalesAdvisorTemplate[salesadvisorID] = salesAdvisorTemplate
				} else if strings.Contains(filename, "竞拍商品") {
					goodsID, err := strconv.ParseInt(row[0], 10, 32)
					if err != nil {
						log.Printf("ERROR: Invalid goodsID in \"%s\" %v", row, err)
						continue
					}
					originalPrice, err := strconv.ParseFloat(row[2], 32)
					if err != nil {
						log.Printf("ERROR: Invalid originalPrice in \"%s\" %v", row, err)
						continue
					}
					limitPrice, err := strconv.ParseFloat(row[3], 32)
					if err != nil {
						log.Printf("ERROR: Invalid limitPrice in \"%s\" %v", row, err)
						continue
					}
					countdownSecond, err := strconv.ParseInt(row[4], 10, 32)
					if err != nil {
						log.Printf("ERROR: Invalid countdownSecond in \"%s\" %v", row, err)
						continue
					}
					auctionGoodsTemplate := map[string]interface{}{
						"goods_name":       row[1],
						"original_price":   float32(originalPrice),
						"limit_price":      float32(limitPrice),
						"countdown_second": int32(countdownSecond),
					}
					AuctionGoodsTemplate[int32(goodsID)] = auctionGoodsTemplate
				}
			}

			log.Printf("Parse config %v ok\n", filename)
		}
	}

	// time.Sleep(time.Second * 3)
	lang, err = json.MarshalIndent(CustomerTemplate, "", "   ")
	if err == nil {
		log.Println(string(lang))
	}
	lang, err = json.MarshalIndent(SalesAdvisorTemplate, "", "   ")
	if err == nil {
		log.Println(string(lang))
	}
	lang, err = json.MarshalIndent(AuctionGoodsTemplate, "", "   ")
	if err == nil {
		log.Println(string(lang))
	}
	log.Printf("Load templated ok.\n\n")
	// time.Sleep(time.Second * 3)

}
