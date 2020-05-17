package config

import (
	"io/ioutil"
	"log"

	"com.lueey.shop/utils"
)

var HosterParams HosterConfig

// HosterConfig 主持人配置
type HosterConfig struct {
	Username string
	Password string
}

// InitHosterConfig 初始化主持人登录参数
// 只有一个，放到 ~/config/ 目录下
func InitHosterConfig() {
	data, err := ioutil.ReadFile(utils.ExpandUser("~/config/host.json"))
	if err != nil {
		log.Fatal("Not found hoster config")
		return
	}
	jst := HosterConfig{}
	log.Printf("Hoster config: %s", data)
	err = json.Unmarshal([]byte(data), &jst)
	if err != nil {
		log.Fatal("The hoster config illegal")
		return
	}

	if jst.Username == "" || jst.Password == "" {
		log.Fatal("Invalid hoster config")
		return
	}

	HosterParams = jst
}
