package main

import (
	"fmt"
	"time"

	"com.lueey.shop/common"
	"com.lueey.shop/config"
	"com.lueey.shop/http/httpHandler"
	"com.lueey.shop/model"
	"com.lueey.shop/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func DummyMiddleware(c *gin.Context) {
	fmt.Println("Im a dummy!")
	c.Next()
}

func main() {
	config.InitDBConfig()
	//
	utils.InitRedisDB()

	common.ServerCategory = common.SERVER_CATEGORY_HTTP

	r := gin.Default()
	store, _ := redis.NewStoreWithDB(10, "tcp", fmt.Sprintf("%s:%s", common.RedisIP, common.RedisPort), common.RedisPass, "2", []byte("secret"))
	store.Options(sessions.Options{MaxAge: utils.HTTPSessionAge})
	r.Use(sessions.Sessions("shop_sessions", store))

	// main.HTTPServerAssignedInit

	r.GET("/time", httpHandler.AssigneeHallserverStep1)

	r.GET("/server_config", httpHandler.AssigneeHallserverStep2)

	r.GET("/uploadfile", httpHandler.AssigneeHallserverStep2)

	r.POST("/host_signin", httpHandler.OnHosterSignin)

	r.POST("/host_update_global_state", httpHandler.OnHosterUpdateState)

	//
	config.Init()

	//
	config.InitHTTPConfig()

	//
	config.InitHosterConfig()

	// 初始化服务器列表
	model.HTTPServerInit()

	// 开始处理未分配服务器的salesAdvisor
	model.HTTPServerAssignedInit()

	// 初始化全球状态
	model.HTTPInitGlobal()

	// crontab 任务
	startTimer(func(now int64) {
		//
		model.HTTPServerDiscovery(now)

		for sid := range model.HTTPServerAllHallServerContainer {
			model.HTTPServerRefresh(now, sid)
		}
	})

	go model.HTTPServerOnServerStartup()
	for sid := range model.HTTPServerAllHallServerContainer {
		go model.HTTPServerOnServerUpdateOnlines(sid)
		go model.HTTPServerOnServerUpdateStatus(sid)
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run(fmt.Sprintf("%s:%s", "0.0.0.0", config.HTTPParams.Port))
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
