package model

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"com.lueey.shop/config"
	"com.lueey.shop/utils"
)

const (
	// ServerActive 启动中
	ServerActive int32 = 1
	// ServerDeactivate 关机
	ServerDeactivate int32 = 0
	// ServerMaxOnline 达到上线
	ServerMaxOnline = 2
	// HTTPServerScanInterval 扫描程序间隔
	HTTPServerScanInterval = 1000
)

var HTTPServerAllHallServerContainer = map[string]*Server{}
var httpServerLastScannerTime int64 = 0
var httpServerLastRefreshTime int64 = 0
var httpServerRwm sync.RWMutex

var TCPServerInstance Server

func GenerateServerKey(sid string) string {
	return fmt.Sprintf("Server##%s", sid)
}

// Server entity
type Server struct {
	ID             string
	status         int32
	HOST           string
	PORT           int32
	online         int32
	limitOnline    int32
	updatedAt      int64
	rwm            sync.RWMutex
	OnlineCustomer []string
}

func (s *Server) String() string {
	lang, err := json.MarshalIndent(s, "", "   ")
	if err == nil {
		return string(lang)
	}
	return ""
}

// TCPServerInit 服务器初始化
func TCPServerInit() {
	cfg := config.InitServerConfig()
	TCPServerInstance = Server{
		ID:          cfg.ServerID,
		HOST:        cfg.ServerHost,
		PORT:        cfg.ServerPort,
		limitOnline: cfg.ServerOnlineLimit,
		status:      ServerActive,
	}

	if v, err := utils.HGetRedis(GenerateServerKey(TCPServerInstance.ID), "status"); err == nil && v == strconv.FormatInt(int64(ServerActive), 10) {
		log.Fatalf("ERROR: The server [%s] has been started.", TCPServerInstance.ID)
	}

	utils.HSetRedis(GenerateServerKey(TCPServerInstance.ID), "id", TCPServerInstance.ID)
	utils.HSetRedis(GenerateServerKey(TCPServerInstance.ID), "host", TCPServerInstance.HOST)
	utils.HSetRedis(GenerateServerKey(TCPServerInstance.ID), "port", TCPServerInstance.PORT)
	utils.HSetRedis(GenerateServerKey(TCPServerInstance.ID), "limitOnline", TCPServerInstance.limitOnline)
	TCPServerInstance.TCPServerUpdateStatus(ServerActive, true)
	TCPServerInstance.TCPServerOnUpdateOnlines(0, true)

	lang, err := json.Marshal(map[string]string{
		"id":          TCPServerInstance.ID,
		"status":      strconv.FormatInt(int64(TCPServerInstance.status), 10),
		"host":        TCPServerInstance.HOST,
		"port":        strconv.FormatInt(int64(TCPServerInstance.PORT), 10),
		"online":      strconv.FormatInt(int64(TCPServerInstance.online), 10),
		"limitOnline": strconv.FormatInt(int64(TCPServerInstance.limitOnline), 10),
		"updatedAt":   strconv.FormatInt(TCPServerInstance.updatedAt, 10),
	})
	if err == nil {
		pubErr := utils.PublishMessage("hallserver##startup", string(lang))
		if pubErr != nil {
			log.Printf("ERROR: publish server startup message failed, because %v", pubErr)
		}
		utils.HSetRedis("hallserver##startup", TCPServerInstance.ID, string(lang))
		log.Printf("New hall server startup and success regist to HTTPServer.\n%s\npublishErr: %v", lang, pubErr)
	} else {
		log.Printf("ERROR: New hall server startup But failed regist to HTTPServer. \n%s\nerr: %v", lang, err)
	}
}

// TCPServerUpdateStatus 更新状态
func (s *Server) TCPServerUpdateStatus(status int32, isInit bool) {
	s.status = status
	utils.HSetRedis(GenerateServerKey(TCPServerInstance.ID), "status", s.status)
	s.updatedAt = utils.NowMilliseconds()
	utils.HSetRedis(GenerateServerKey(TCPServerInstance.ID), "updatedAt", s.updatedAt)
	log.Printf("Notify HTTPServer status %d by serverKey %v", s.status, s.ID)

	if !isInit {
		utils.PublishMessage(fmt.Sprintf("%s##status", GenerateServerKey(s.ID)), strconv.FormatInt(int64(s.status), 10))
	}
}

// TCPServerOnUpdateOnlines 更新服务器在线人数
func (s *Server) TCPServerOnUpdateOnlines(onlines int32, isInit bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	s.online = onlines

	if s.online >= s.limitOnline { // 上升到最大人数
		s.TCPServerUpdateStatus(ServerMaxOnline, false)
	}

	utils.HSetRedis(GenerateServerKey(s.ID), "online", s.online)
	s.updatedAt = utils.NowMilliseconds()
	utils.HSetRedis(GenerateServerKey(TCPServerInstance.ID), "updatedAt", s.updatedAt)
	log.Printf("Notify HTTPServer onlines %d by serverKey %v", s.online, s.ID)
	if !isInit {
		utils.PublishMessage(fmt.Sprintf("%s##onlines", GenerateServerKey(s.ID)), strconv.FormatInt(int64(s.online), 10))
	}
}

// TCPServerRefresh delay update
func (s *Server) TCPServerRefresh(now int64) {
	if now-s.updatedAt < 1000 {
		return
	}
	s.updatedAt = utils.NowMilliseconds()
	utils.HSetRedis(GenerateServerKey(TCPServerInstance.ID), "updatedAt", s.updatedAt)
}

// HTTPServerInit 依据redis数据初始化大厅服务列表
func HTTPServerInit() {
	log.Println("Start initialize hall server list from redis data.")
	if m, err := utils.HGetAllRedis("hallserver##startup"); err == nil {
		for hallServerID := range m {
			if _, ok := HTTPServerAllHallServerContainer[hallServerID]; !ok {
				swapMap := map[string]string{}
				if id, err := utils.HGetRedis(GenerateServerKey(hallServerID), "id"); err == nil {
					swapMap["id"] = id.(string)
				} else {
					log.Printf("Can not parse server obj from redis data, because has not field 'ID'")
					continue
				}
				if host, err := utils.HGetRedis(GenerateServerKey(hallServerID), "host"); err == nil {
					swapMap["host"] = host.(string)
				} else {
					log.Printf("Can not parse server obj from redis data, because has not field 'HOST'")
					continue
				}
				if port, err := utils.HGetRedis(GenerateServerKey(hallServerID), "port"); err == nil {
					swapMap["port"] = port.(string)
				} else {
					log.Printf("Can not parse server obj from redis data, because has not field 'PORT'")
					continue
				}
				if limitOnline, err := utils.HGetRedis(GenerateServerKey(hallServerID), "limitOnline"); err == nil {
					swapMap["limitOnline"] = limitOnline.(string)
				} else {
					log.Printf("Can not parse server obj from redis data, because has not field 'limitOnline'")
					continue
				}
				if status, err := utils.HGetRedis(GenerateServerKey(hallServerID), "status"); err == nil {
					swapMap["status"] = status.(string)
				} else {
					log.Printf("Can not parse server obj from redis data, because has not field 'status'")
					continue
				}
				if online, err := utils.HGetRedis(GenerateServerKey(hallServerID), "online"); err == nil {
					swapMap["online"] = online.(string)
				} else {
					log.Printf("Can not parse server obj from redis data, because has not field 'online'")
					continue
				}
				if updatedAt, err := utils.HGetRedis(GenerateServerKey(hallServerID), "updatedAt"); err == nil {
					swapMap["updatedAt"] = updatedAt.(string)
				} else {
					log.Printf("Can not parse server obj from redis data, because has not field 'updatedAt'")
					continue
				}

				s := parseServerFromMap(swapMap)
				if s != nil {
					HTTPServerAllHallServerContainer[s.ID] = s
				}
			}
		}
	}
}

// HTTPServerAssignedInit 给所有房间分配服务器
func HTTPServerAssignedInit() {
	for salesAdvisorID := range config.SalesAdvisorTemplate {
		SelectHallServer(salesAdvisorID, false)
	}
}

func parseServerFromMap(a map[string]string) *Server {
	log.Printf("Will be create new hall server by %v.", a)
	s := Server{}
	if id, ok := a["id"]; ok {
		s.ID = id
	} else {
		log.Println("ERROR: But the hall server has no ID!")
		return nil
	}
	if status, ok := a["status"]; ok {
		if statusInt, err := strconv.ParseInt(status, 10, 32); err == nil {
			s.status = int32(statusInt)
		} else {
			log.Println("ERROR: But the hall server status illegal.", status)
			return nil
		}
	} else {
		log.Printf("ERROR: But the hall server has no status!")
		return nil
	}
	if host, ok := a["host"]; ok {
		s.HOST = host
	} else {
		log.Println("ERROR: But the hall server has no host!")
		return nil
	}
	if port, ok := a["port"]; ok {
		if portInt, err := strconv.ParseInt(port, 10, 32); err == nil {
			s.PORT = int32(portInt)
		} else {
			log.Println("ERROR: But the hall server port illegal.", port)
			return nil
		}
	} else {
		log.Println("ERROR: But the hall server has no port!")
		return nil
	}
	if online, ok := a["online"]; ok {
		if onlineInt, err := strconv.ParseInt(online, 10, 32); err == nil {
			s.online = int32(onlineInt)
		} else {
			log.Println("ERROR: But the hall server online illegal.", online)
			return nil
		}
	} else {
		log.Printf("ERROR: But the hall server has no online!")
		return nil
	}
	if limitOnline, ok := a["limitOnline"]; ok {
		if limitOnlineInt, err := strconv.ParseInt(limitOnline, 10, 32); err == nil {
			s.limitOnline = int32(limitOnlineInt)
		} else {
			log.Println("ERROR: But the hall server limitOnline illegal.", limitOnline)
			return nil
		}
	} else {
		log.Printf("ERROR: But the hall server has no limitOnline!")
		return nil
	}
	if updatedAt, ok := a["updatedAt"]; ok {
		if updatedAtLong, err := strconv.ParseInt(updatedAt, 10, 64); err == nil {
			s.updatedAt = updatedAtLong
		} else {
			log.Println("ERROR: But the hall server updatedAt illegal.", updatedAt)
			return nil
		}
	} else {
		log.Printf("ERROR: But the hall server has no updatedAt!")
		return nil
	}
	return &s
}

// HTTPServerOnServerStartup 功能服务器启动
func HTTPServerOnServerStartup() {
	pubsub := utils.GetRDB().Subscribe("hallserver##startup")
	defer func() { pubsub.Close() }()

	ch := pubsub.ChannelSize(1)
	for {
		res := <-ch
		log.Printf("INFO: recived hallserver OnStartup message: %s", res.Payload)
		a := map[string]string{}
		if err := json.Unmarshal([]byte(res.Payload), &a); err == nil {
			s := parseServerFromMap(a)
			if s == nil {
				continue
			}
			if s.status != ServerActive {
				log.Printf("ERROR: discovered new hall server, But status is not 'Active'. %v", s)
				continue
			}
			HTTPServerAllHallServerContainer[s.ID] = s
		} else {
			log.Printf("ERROR: HTTPSERVER discovered new hall server, But the hall server data illegal. %v. %v", a, err)
		}
	}
}

// HTTPServerOnServerUpdateOnlines 更新服务器在线人数
func HTTPServerOnServerUpdateOnlines(sid string) {
	pubsub := utils.GetRDB().Subscribe(fmt.Sprintf("%s##onlines", GenerateServerKey(sid)))
	defer func() { pubsub.Close() }()

	ch := pubsub.ChannelSize(1)
	for {
		res := <-ch
		s, ok := HTTPServerAllHallServerContainer[sid]
		if !ok {
			continue
		}
		log.Printf("INFO: recived hallserver UpdateOnlines message: %s", res.Payload)
		if v, err := strconv.ParseInt(res.Payload, 10, 32); err == nil {
			s.online = int32(v)
			log.Printf("INFO: Server[%s] update onlines to %d.", GenerateServerKey(s.ID), s.online)
		} else {
			log.Println("ERROR: ", err)
		}
	}
}

// HTTPServerOnServerUpdateStatus 更新服务器状态
func HTTPServerOnServerUpdateStatus(sid string) {
	pubsub := utils.GetRDB().Subscribe(fmt.Sprintf("%s##status", GenerateServerKey(sid)))
	defer func() { pubsub.Close() }()

	ch := pubsub.ChannelSize(1)
	for {
		res := <-ch
		s, ok := HTTPServerAllHallServerContainer[sid]
		if !ok {
			continue
		}
		log.Printf("INFO: recived hallserver UpdatesStatus message: %s", res.Payload)
		if v, err := strconv.ParseInt(res.Payload, 10, 32); err == nil {
			s.status = int32(v)
			log.Printf("INFO: Server[%s] update status to %d.", GenerateServerKey(s.ID), s.status)
		} else {
			log.Println("ERROR: ", err)
		}
	}
}

// HTTPServerDiscovery 发现新的服务器
func HTTPServerDiscovery(now int64) {
	if now-httpServerLastScannerTime < HTTPServerScanInterval {
		return
	}
	httpServerLastScannerTime = now
	// if config.DEBUG {
	// 	log.Println("Start scan and try get new hall server.")
	// }
	if m, err := utils.HGetAllRedis("hallserver##startup"); err == nil {
		for hallServerID, serverJsonstr := range m {
			if _, ok := HTTPServerAllHallServerContainer[hallServerID]; !ok {
				a := map[string]string{}
				if err := json.Unmarshal([]byte(serverJsonstr), &a); err == nil {
					s := parseServerFromMap(a)
					if s == nil {
						continue
					}
					if s.status == ServerDeactivate {
						log.Printf("ERROR: scan new hall server, But status is Deactive. %v", s)
						continue
					}
					HTTPServerAllHallServerContainer[s.ID] = s
				} else {
					log.Printf("ERROR: HTTPSERVER discovered new hall server, But the hall server data illegal. %v. %v", a, err)
				}
			}
		}
	}
}

// HTTPServerRefresh 定时刷新服务器信息
func HTTPServerRefresh(now int64, sid string) {
	if now-httpServerLastRefreshTime < HTTPServerScanInterval {
		return
	}
	s, ok := HTTPServerAllHallServerContainer[sid]
	if !ok {
		return
	}
	httpServerLastRefreshTime = now
	var lastHeartbeat int64 = 0
	if v, err := utils.HGetRedis(GenerateServerKey(s.ID), "updatedAt"); err == nil {
		if z, err := strconv.ParseInt(v.(string), 10, 64); err == nil {
			lastHeartbeat = z
		} else {
			log.Println("ERROR: ", err)
		}
	} else {
		log.Println("ERROR: ", err)
	}

	if lastHeartbeat == 0 {
		log.Printf("ERROR: Invalid server updatedAt value, close server entry.")
		s.status = ServerDeactivate
		return
	}

	if now-lastHeartbeat >= 5000 { // 5秒内未更新则认为服务器已宕机
		s.status = ServerDeactivate
		log.Printf("ERROR: Server[%s] without vital signs. No heartbeat detected.", s.ID)
		return
	}
	if s.status == ServerDeactivate {
		s.status = ServerActive
		log.Printf("WARN: Server[%s] resurrection.", s.ID)
	}

	if v, err := utils.HGetRedis(GenerateServerKey(s.ID), "online"); err == nil {
		if z, err := strconv.ParseInt(v.(string), 10, 32); err == nil {
			s.limitOnline = int32(z)
		} else {
			log.Println("ERROR: ", err)
		}
	}
	s.updatedAt = lastHeartbeat
}

// HTTPServerGetServerAddress 返回当前服务器地址
func (s *Server) HTTPServerGetServerAddress() string {
	return fmt.Sprintf("%v:%d", s.HOST, s.PORT)
}

// SelectHallServer 找到销售人员对应的服务器，如果没有结果，则对其分配一个最优服务器并返回
func SelectHallServer(salesAdvisorID string, prowlNotify bool) string {
	httpServerRwm.Lock()
	defer httpServerRwm.Unlock()

	assignedHallServerID, err := utils.HGetRedis("SalesAdvisor###ServerID", salesAdvisorID)
	assigned := err == nil && assignedHallServerID != "" // 是否已经分配过server
	var maxOnlinesServer *Server = nil
	if assigned { // 是否需要重新分配
		s, ok := HTTPServerAllHallServerContainer[assignedHallServerID.(string)]
		if !ok {
			log.Printf("ERROR: hall serverID is invalid and cannot used assigned server instance! %v %s", s.ID)
			return ""
			// assigned = false
			// } else if s.status == ServerDeactivate {
			// 	log.Printf("ERROR: The sales roome server has been shut down And needs to be reassigned. %v", s.ID)
			// 	assigned = false
		}
	}
	if !assigned { // 还没有分配服务器
		for _, s := range HTTPServerAllHallServerContainer {
			if s.status == ServerDeactivate || s.status == ServerMaxOnline {
				continue
			}
			if maxOnlinesServer == nil || maxOnlinesServer.online < s.online { // 一个满了再去另一个（这个过程可以不断调整上限）
				maxOnlinesServer = s
			}
		}
		if maxOnlinesServer == nil {
			errorMsg := fmt.Sprintf("Can not found available hall server assigned to sales advisor. salesID[%s]", salesAdvisorID)
			log.Println("ERROR: ", errorMsg)
			if prowlNotify {
				utils.ProwlNotify(errorMsg)
			}
			return ""
		}
		utils.HSetRedis("SalesAdvisor###ServerID", salesAdvisorID, maxOnlinesServer.ID)
		return maxOnlinesServer.HTTPServerGetServerAddress()

	}
	return HTTPServerAllHallServerContainer[assignedHallServerID.(string)].HTTPServerGetServerAddress()
}
