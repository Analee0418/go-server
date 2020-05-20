package config

import (
	"io/ioutil"
	"log"

	"com.lueey.shop/common"
	"com.lueey.shop/utils"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// DBConfig 数据库连接配置
type DBConfig struct {
	DBhost string
	DBport string
	DBuser string
	DBpass string
}

// RedisConfig Redis连接配置
type RedisConfig struct {
	RedisIP   string
	RedisPort string
	RedisPass string
}

// InitDBConfig 初始化主持人登录参数
// 只有一个，放到 ~/config/ 目录下
func InitDBConfig() {
	// DB Config
	data, err := ioutil.ReadFile(utils.ExpandUser("~/config/db.json"))
	if err != nil {
		log.Fatal("FATAL: Not found DB config")
		return
	}
	dbJSON := DBConfig{}
	log.Printf("[INFO] DB config: %s", data)
	err = json.Unmarshal([]byte(data), &dbJSON)
	if err != nil {
		log.Fatal("FATAL: The DB config illegal")
		return
	}

	if dbJSON.DBhost == "" || dbJSON.DBport == "" {
		log.Printf("\033[1;33m[WARNING] \033[0mInvalid DB config, trying to used default params.")
		dbJSON.DBhost = "127.0.0.1"
		dbJSON.DBport = "27017"
	}
	common.DBHOST = dbJSON.DBhost
	common.DBPORT = dbJSON.DBport
	common.DBUSER = dbJSON.DBuser
	common.DBPASS = dbJSON.DBpass

	// Redis Config
	data, err = ioutil.ReadFile(utils.ExpandUser("~/config/redis.json"))
	if err != nil {
		log.Fatal("FATAL: Not found Redis config")
		return
	}
	redisJSON := RedisConfig{}
	log.Printf("[INFO] Redis config: %s", data)
	err = json.Unmarshal([]byte(data), &redisJSON)
	if err != nil {
		log.Fatal("FATAL: The Redis config illegal")
		return
	}

	if redisJSON.RedisIP == "" || redisJSON.RedisPort == "" {
		log.Printf("\033[1;33m[WARNING] \033[0mInvalid Redis config, trying to used default params.")
		redisJSON.RedisIP = "127.0.0.1"
		redisJSON.RedisPort = "27017"
	}
	common.RedisIP = redisJSON.RedisIP
	common.RedisPort = redisJSON.RedisPort
	common.RedisPass = redisJSON.RedisPass
}
