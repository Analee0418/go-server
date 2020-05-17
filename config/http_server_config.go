package config

import (
	"io/ioutil"
	"log"

	"com.lueey.shop/utils"
)

var HTTPParams HTTPConfig

// HTTPConfig 主持人配置
type HTTPConfig struct {
	IP            string
	Port          string
	RequireHeader string
}

// InitHTTPConfig 初始化主持人登录参数
// 只有一个，放到 ~/config/ 目录下
func InitHTTPConfig() {
	data, err := ioutil.ReadFile(utils.ExpandUser("~/config/http.json"))
	if err != nil {
		log.Print("WARNING: Not found HTTP config")
		HTTPParams = HTTPConfig{IP: "http://127.0.0.1", Port: "11001"}
		return
	}
	jst := HTTPConfig{}
	log.Printf("HTTP config: %s", data)
	err = json.Unmarshal([]byte(data), &jst)
	if err != nil {
		log.Fatal("The HTTP config illegal")
		return
	}

	if jst.IP == "" || jst.Port == "" {
		log.Fatal("Invalid HTTP config")
		return
	}

	HTTPParams = jst
}
