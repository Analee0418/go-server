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
		log.Fatal("Not found DB config")
		return
	}
	jst := DBConfig{}
	log.Printf("DB config: %s", data)
	err = json.Unmarshal([]byte(data), &jst)
	if err != nil {
		log.Fatal("The DB config illegal")
		return
	}

	log.Println(jst)

	if jst.DBhost == "" || jst.DBport == "" {
		log.Printf("WARNING: Invalid DB config, trying to used default params.")
		jst.DBhost = "127.0.0.1"
		jst.DBport = "27017"
	}
	common.DBHOST = jst.DBhost
	common.DBPORT = jst.DBport
	common.DBUSER = jst.DBuser
	common.DBPASS = jst.DBpass

	// Redis Config
	data, err = ioutil.ReadFile(utils.ExpandUser("~/config/redis.json"))
	if err != nil {
		log.Fatal("Not found Redis config")
		return
	}
	jst2 := RedisConfig{}
	log.Printf("Redis config: %s", data)
	err = json.Unmarshal([]byte(data), &jst2)
	if err != nil {
		log.Fatal("The Redis config illegal")
		return
	}

	log.Println(jst2)

	if jst2.RedisIP == "" || jst2.RedisPort == "" {
		log.Printf("WARNING: Invalid Redis config, trying to used default params.")
		jst2.RedisIP = "127.0.0.1"
		jst2.RedisPort = "27017"
	}
	common.RedisIP = jst2.RedisIP
	common.RedisPort = jst2.RedisPort
	common.RedisPass = jst2.RedisPass
}
