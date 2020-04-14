package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"com.lueey.shop/utils"
)

const DEBUG bool = true

// ServerConfig 定义配置文件解析后的结构
type ServerConfig struct {
	ServerID          string
	ServerHost        string
	ServerPort        int32
	ServerOnlineLimit int32
}

func InitServerConfig() (jst *ServerConfig) {
	data, err := ioutil.ReadFile(utils.ExpandUser("./server.json"))
	if err != nil {
		log.Fatal("Not found server config")
		return
	}
	jst = &ServerConfig{}
	log.Printf("Server config: %s", data)
	err = json.Unmarshal([]byte(data), jst)
	if err != nil {
		log.Fatal("The server config illegal")
		return
	}

	if jst.ServerHost == "" || jst.ServerPort == 0 || jst.ServerID == "" {
		log.Fatal("Invalid server config")
		return
	}

	return jst
}
