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
		log.Fatal("FATAL: Not found hoster config")
		return
	}
	hostJSON := HosterConfig{}
	log.Printf("[INFO] Hoster config: %s", data)
	err = json.Unmarshal([]byte(data), &hostJSON)
	if err != nil {
		log.Fatal("FATAL: The hoster config illegal")
		return
	}

	if hostJSON.Username == "" || hostJSON.Password == "" {
		log.Fatal("[INFO] Invalid hoster config")
		return
	}

	HosterParams = hostJSON
}
