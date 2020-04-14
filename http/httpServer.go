package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
	"time"

	"com.lueey.shop/config"
	"com.lueey.shop/model"
	"com.lueey.shop/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/time", func(c *gin.Context) {
		// simulate a long task with time.Sleep(). 5 seconds

		if v, e := c.Cookie("H_skeys_identity"); v != "87fb614a3b04e9a01fcdf" || e != nil {
			c.String(403, "Invalid request.")
			return
		}
		z := time.Now().String()
		c.SetCookie("jetztMal", string(z), 600, "/", "*", false, true)
		c.String(200, "ok")
	})

	r.GET("/server_config", func(c *gin.Context) {
		// simulate a long task with time.Sleep(). 5 seconds

		log.Println(c.Request.Cookies())
		ts, exists := c.GetQuery("t")
		if !exists {
			log.Println("ERROR: Not found jetztMal")
			c.String(403, "Invalid request.")
			return
		}
		var salesAdvisorID string
		salesAdvisorID, exists = c.GetQuery("s")
		if !exists || salesAdvisorID == "" {
			idcard, exists := c.GetQuery("c")
			if !exists {
				log.Println("ERROR: Not found users params")
				c.String(403, "Invalid request.")
				return
			}
			if v, ok := config.CustomerTemplate[idcard]; ok {
				salesAdvisorID = v["sales_advisor"]
			} else {
				log.Printf("Can not found customer config by idcard[%s]", idcard)
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
	})

	//
	config.Init()

	// 初始话服务器列表
	model.HTTPServerInit()

	// 开始处理未分配服务器的salesAdvisor
	model.HTTPServerAssignedInit()

	// crontab 任务
	startTimer(func(now int64) {
		//
		model.HTTPServerDiscovery(now)

		for _, server := range model.HTTPServerAllHallServerContainer {
			server.HTTPServerRefresh(now)
		}
	})

	go model.HTTPServerOnServerStartup()
	for _, server := range model.HTTPServerAllHallServerContainer {
		go server.HTTPServerOnServerUpdateOnlines()
		go server.HTTPServerOnServerUpdateStatus()
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run(":11001")
}

func startTimer(f func(int64)) {
	go func() {
		for {
			now := time.Now()
			f(utils.NowMillisecondsByTime(now))
			next := now.Add(time.Millisecond * 100)
			next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), next.Second(), next.Nanosecond(), next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}
