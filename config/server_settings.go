package config

import (
	"io/ioutil"
	"log"

	"com.lueey.shop/utils"
)

const DEBUG bool = true

var TCPServerParams ServerConfig

// ServerConfig 定义配置文件解析后的结构
// 配置与可执行文件放到一起
type ServerConfig struct {
	ServerID          string
	ServerHost        string
	ServerPort        int32
	ServerOnlineLimit int32
}

func InitServerConfig() (jst *ServerConfig) {
	data, err := ioutil.ReadFile(utils.ExpandUser("./server.json"))
	if err != nil {
		log.Fatal("FATAL: Not found server config")
		return
	}
	jst = &ServerConfig{}
	log.Printf("INFO: Server config: %s", data)
	err = json.Unmarshal([]byte(data), jst)
	if err != nil {
		log.Fatal("FATAL: The server config illegal")
		return
	}

	if jst.ServerHost == "" || jst.ServerPort == 0 || jst.ServerID == "" {
		log.Fatal("FATAL: Invalid server config")
		return
	}

	TCPServerParams = *jst
	return jst
}
