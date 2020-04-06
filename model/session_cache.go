package model

import (
	"log"
	"net"

	"com.lueey.shop/config"
	"com.lueey.shop/utils"
	guuid "github.com/google/uuid"
)

var lastReleaseTime int64 = utils.NowMilliseconds()

var sessionCacheByUserName = map[string]*Session{}
var sessionCacheByID = map[guuid.UUID]*Session{}
var sessionConn = map[net.Conn]*Session{}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
}

func SessionByID(uuid string) (*Session, bool) {
	if len(uuid) == 0 {
		return nil, false
	}
	sid, e := guuid.Parse(uuid)
	if e != nil {
		log.Printf("Error: %v", e)
		return nil, false
	}
	if s, ok := sessionCacheByID[sid]; ok {
		return s, true
	}
	return nil, false
}

func GetSessionByName(name string) (*Session, bool) {
	if s, ok := sessionCacheByUserName[name]; ok {
		return s, true
	}
	return nil, false
}

func AddSession(conn net.Conn, s *Session) {
	if _, ok := sessionCacheByID[s.id]; ok {
		log.Printf("has been exists in cache: %v\n", s)
	} else {
		sessionCacheByID[s.id] = s
		sessionCacheByUserName[s.name] = s
		sessionConn[conn] = s
		log.Printf("add new session to cache: %v\n", s)
	}
}

func DeleteSession(uuid guuid.UUID, name string, conn net.Conn) {
	defer func() { conn.Close() }()
	if _, ok := sessionCacheByID[uuid]; ok {
		delete(sessionCacheByID, uuid)
		log.Printf("has been deleted session from cache_ID: %v\n", uuid)
	}

	if _, ok := sessionCacheByUserName[name]; ok {
		delete(sessionCacheByUserName, name)
		log.Printf("has been deleted session from cache_Name: %v\n", name)
	}

	if _, ok := sessionConn[conn]; ok {
		delete(sessionConn, conn)
		log.Printf("has been deleted session from cache_Conn: %v\n", conn)
	}

	log.Printf("DEBUG: session.length: id.%d, name.%d, conn.%d", len(sessionCacheByID), len(sessionCacheByUserName), len(sessionConn))
}

func ReleaseSessionCache(now int64) {
	if now-lastReleaseTime < 1000 {
		return
	}

	// log.Printf("Start finding invalid session. %d\n", len(sessionCacheByID))

	invalidKeys := []func(){}
	for _, session := range sessionCacheByID {
		if session.dead {
			invalidKeys = append(invalidKeys, func() { DeleteSession(session.id, session.name, session.conn) })
		} else {
			if utils.NowMilliseconds()-session.lastHeartBeatMillisecond > func() int64 {
				if config.DEBUG {
					return 60000
				} else {
					return 30000
				}
			}() {
				conn := session.conn
				invalidKeys = append(invalidKeys, func() { DeleteSession(session.id, session.name, conn) })
				session.Close("Connection without vital signs. No heartbeat detected.")

				// 放置 goroutine queue
				// go DeleteSession(conn, uuid, name)
			}
		}
	}

	// log.Printf("Invalid sesisons count %d", len(invalidKeys))
	for _, fun := range invalidKeys {
		fun()
	}

}

func OnClientDisconnect(conn net.Conn) {
	defer func() { conn.Close() }()
	s, ok := sessionConn[conn]
	if ok {
		if s.customerInfo != nil {
			GlobalOnCustomerDisconnect(s.customerInfo.ID)
			s.customerInfo = nil
		}
		s.conn.Close()
		s.dead = true
		log.Printf("Client disconnect after clear session[%s].", s.name)
	}

}
