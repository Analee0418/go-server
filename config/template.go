package config

import (
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"com.lueey.shop/common"
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
	log.Println("INFO: Start to load templated.")
	files, _ := filepath.Glob(utils.ExpandUser("~/config/*.xlsx"))
	lang, err := json.MarshalIndent(files, "", "   ")
	if err == nil {
		strs := string(lang)
		log.Printf("INFO: ==================== ALL CONFIG FILES %s", strs) // contains a list of all files in the current directory

		for _, filename := range files {
			log.Printf("INFO: Parse config %v", filename)
			xlsx, err := excelize.OpenFile(filename)
			if err != nil {
				log.Panic("PANIC: ", err)
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
				log.Printf("INFO: %s", strings.Join(rows[0], ",\t"))
			}
			lastAdvisorName := ""
			lastAdvisorMobile := ""
			lastAdvisorID := ""
			lastProvince := ""
			lastCity := ""
			lastCompany := ""
			// 销售手机号	销售身份证号	经商销省份	经商销城市	经商销公司名称
			for _, row := range rows[1:] {
				if DEBUG {
					log.Printf("INFO: %s", strings.Join(row, ",\t"))
				}

				if strings.Contains(filename, "销售") {
					// 销售及用户的配置表
					idcard := row[0]
					mobile := row[1]
					region := ""

					if row[4] != "" {
						lastAdvisorName = row[4]
					}
					if row[5] != "" {
						lastAdvisorMobile = row[5]
					}
					if row[6] != "" {
						lastAdvisorID = row[6]
					}
					if row[7] != "" {
						lastProvince = row[7]
					}
					if row[8] != "" {
						lastCity = row[8]
					}
					if row[9] != "" {
						lastCompany = row[9]
					}

					if val, err := utils.HGetRedis("mobileRegion", mobile); err == nil {

						region = val.(string)
						log.Printf("INFO: Mobile region from redis: %s", region)

					} else if common.ServerCategory == common.SERVER_CATEGORY_HALL {

						resp, err := http.Get("https://www.ip.cn/db?num=%2B" + mobile)
						if err != nil {
							log.Printf("\033[1;33mWARNING: \033[0mcrawl mobile region failed. %v", err)
							continue
						}

						if resp.StatusCode == 200 {
							body, err := ioutil.ReadAll(resp.Body)
							if err != nil {
								log.Printf("\033[1;33mWARNING: \033[0mread mobile region failed. %v", err)
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
						"sales_advisor": lastAdvisorID,
					}
					CustomerTemplate[idcard] = customerTemplate

					// salesadvisorID := row[5]
					salesAdvisorTemplate := map[string]string{
						"advisor_name":     lastAdvisorName,
						"advisor_id":       lastAdvisorID,
						"advisor_mobile":   lastAdvisorMobile,
						"advisor_province": lastProvince,
						"advisor_city":     lastCity,
						"advisor_company":  lastCompany,
					}
					SalesAdvisorTemplate[lastAdvisorID] = salesAdvisorTemplate

				} else if strings.Contains(filename, "竞拍商品") {
					// 竞拍的配置表
					goodsID, err := strconv.ParseInt(row[0], 10, 32)
					if err != nil {
						log.Printf("\033[1;33mWARNING: \033[0mInvalid goodsID in \"%s\" %v", row, err)
						continue
					}
					originalPrice, err := strconv.ParseFloat(row[2], 32)
					if err != nil {
						log.Printf("\033[1;33mWARNING: \033[0mInvalid originalPrice in \"%s\" %v", row, err)
						continue
					}
					limitPrice, err := strconv.ParseFloat(row[3], 32)
					if err != nil {
						log.Printf("\033[1;33mWARNING: \033[0mInvalid limitPrice in \"%s\" %v", row, err)
						continue
					}
					countdownSecond, err := strconv.ParseInt(row[4], 10, 32)
					if err != nil {
						log.Printf("\033[1;33mWARNING: \033[0mInvalid countdownSecond in \"%s\" %v", row, err)
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

			log.Printf("INFO: Parse config %v ok\n", filename)
		}
	}

	// time.Sleep(time.Second * 3)
	lang, err = json.MarshalIndent(CustomerTemplate, "", "   ")
	if err == nil {
		log.Println("INFO: ", string(lang))
	}
	lang, err = json.MarshalIndent(SalesAdvisorTemplate, "", "   ")
	if err == nil {
		log.Println("INFO: ", string(lang))
	}
	// lang, err = json.MarshalIndent(AuctionGoodsTemplate, "", "   ")
	// if err == nil {
	// 	log.Println("INFO: ", string(lang))
	// }
	log.Printf("INFO: Load templated ok.\n\n")
	// time.Sleep(time.Second * 3)

}
