package httpHandler

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
	"time"

	"com.lueey.shop/config"
	"com.lueey.shop/model"
	"github.com/gin-gonic/gin"
)

func AssigneeHallserverStep1(c *gin.Context) {
	// simulate a long task with time.Sleep(). 5 seconds

	if v, e := c.Cookie("H_skeys_identity"); v != "87fb614a3b04e9a01fcdf" || e != nil {
		c.String(403, "Invalid request.")
		return
	}
	z := time.Now().String()
	c.SetCookie("jetztMal", string(z), 600, "/", "*", false, true)
	c.String(200, "ok")
}

func AssigneeHallserverStep2(c *gin.Context) {
	// simulate a long task with time.Sleep(). 5 seconds

	log.Println(c.Request.Cookies())
	ts, exists := c.GetQuery("t")
	if !exists {
		log.Println("ERROR: Not found jetztMal")
		c.String(403, "Invalid request.")
		return
	}
	var salesAdvisorID string
	salesAdvisorID, exists = c.GetQuery("s") // 销售
	if !exists || salesAdvisorID == "" {
		idcard, exists := c.GetQuery("c") // 用户ID
		if !exists {
			log.Println("ERROR: Not found users ID")
			c.String(403, "Invalid request.")
			return
		}
		mobile, exists := c.GetQuery("m") // 用户手机号
		if !exists {
			log.Println("ERROR: Not found users mobile")
			c.String(403, "Invalid request.")
			return
		}

		if v, ok := config.CustomerTemplate[idcard]; !ok {
			for templateID, template := range config.CustomerTemplate {
				if l := len(templateID); templateID[l-6:l] == idcard && template["mobile"] == mobile {
					idcard = templateID
					salesAdvisorID = template["sales_advisor"]
					break
				}
			}
		} else {
			salesAdvisorID = v["sales_advisor"]
		}

		if salesAdvisorID == "" {
			log.Printf("Can not found customer config by idcard[%s] mobile[%s]", idcard, mobile)
		}
	}

	if config.DEBUG {
		log.Printf("SalesAdvisroID: %s, idCard: %v", salesAdvisorID, c)
	}

	if _, ok := config.SalesAdvisorTemplate[salesAdvisorID]; !ok {
		log.Println("ERROR: Invalid user params", c.Params)
		lang, err := json.MarshalIndent(c.Params, "", "   ")
		if err == nil {
			log.Println(string(lang))
		}
		c.String(403, "Invalid request.")
		return
	}

	v, e := c.Cookie("shop_SID")
	if e != nil {
		log.Println("ERROR: Not found shop_SID to used verify")
		c.String(403, "Invalid request.")
		return
	}

	hasher := md5.New()
	hasher.Write([]byte(ts + "5d56179ecd32148eec0021178b9b2e83"))
	if v != hex.EncodeToString(hasher.Sum(nil)) {
		log.Printf("invalid sid, remote: %s, local: %s", v, hex.EncodeToString(hasher.Sum(nil)))
		c.String(403, "Invalid request.")
		return
	}

	result := model.SelectHallServer(salesAdvisorID, true)
	c.String(200, result)
}
