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
		log.Print("\033[1;33mWARNING: \033[0mNot found HTTP config, used DEFAULT config 127.0.0.1:11001")
		HTTPParams = HTTPConfig{IP: "http://127.0.0.1", Port: "11001"}
		return
	}
	httpJSON := HTTPConfig{}
	log.Printf("INFO: HTTP config: %s", data)
	err = json.Unmarshal([]byte(data), &httpJSON)
	if err != nil {
		log.Fatal("FATAL: The HTTP config illegal")
		return
	}

	if httpJSON.IP == "" || httpJSON.Port == "" {
		log.Fatal("FATAL: Invalid HTTP config")
		return
	}

	HTTPParams = httpJSON
}
