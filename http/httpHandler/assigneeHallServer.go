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
		c.String(403, "无效请求")
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
		c.String(403, "无效请求")
		return
	}

	v, e := c.Cookie("shop_SID")
	if e != nil {
		log.Println("ERROR: Not found shop_SID to used verify")
		c.String(403, "无效请求")
		return
	}

	hasher := md5.New()
	hasher.Write([]byte(ts + "5d56179ecd32148eec0021178b9b2e83"))
	if v != hex.EncodeToString(hasher.Sum(nil)) {
		log.Printf("ERROR: invalid sid, remote: %s, local: %s", v, hex.EncodeToString(hasher.Sum(nil)))
		c.String(403, "无效请求")
		return
	}

	var salesAdvisorID string = ""

	mobile, exists := c.GetQuery("m") // 用户/销售 手机号
	if !exists {
		log.Println("ERROR: Not found parameter mobile")
		c.String(403, "无效请求")
		return
	}

	if salesAdvisorID, exists = c.GetQuery("s"); exists {
		// 销售
		for templateID, template := range config.SalesAdvisorTemplate {
			if l := len(templateID); templateID[l-6:l] == salesAdvisorID && template["advisor_mobile"] == mobile {
				salesAdvisorID = templateID
				break
			}
		}
		if salesAdvisorID == "" {
			log.Printf("ERROR: Can not found sales advisor config by idcard[%s] mobile[%s]", salesAdvisorID, mobile)
			c.String(403, "销售ID或手机号错误.")
			return
		}
	} else if customerID, exists := c.GetQuery("c"); exists {
		// 用户ID
		for templateID, template := range config.CustomerTemplate {
			if l := len(templateID); templateID[l-6:l] == customerID && template["mobile"] == mobile {
				customerID = templateID
				salesAdvisorID = template["sales_advisor"]
				break
			}
		}
		if salesAdvisorID == "" {
			log.Printf("ERROR: Can not found customer config by idcard[%s] mobile[%s]", customerID, mobile)
			c.String(403, "用户ID或手机号错误")
			return
		}
	} else {
		log.Println("ERROR: Not found customerID or salesAdvisorID")
		c.String(403, "参数错误")
		return
	}

	if config.DEBUG {
		log.Printf("Final SalesAdvisroID: %s", salesAdvisorID)
	}

	if _, ok := config.SalesAdvisorTemplate[salesAdvisorID]; !ok {
		log.Println("ERROR: Invalid user params", c.Params)
		lang, err := json.MarshalIndent(c.Params, "", "   ")
		if err == nil {
			log.Println(string(lang))
		}
		c.String(403, "无效请求")
		return
	}

	result := model.SelectHallServer(salesAdvisorID, true)
	c.String(200, result)
}
