package model

import (
	"log"

	"com.lueey.shop/utils"
	guuid "github.com/google/uuid"
)

var lastReleaseTime int64 = utils.NowMilliseconds()

var sessionCacheByUserName = map[string]*Session{}
var sessionCacheByID = map[guuid.UUID]*Session{}

func SessionByID(uuid string) (*Session, bool) {
	log.Println(`+++++++++++++++++++++`)
	if len(uuid) == 0 {
		return nil, false
	}
	sid, e := guuid.Parse(uuid)
	if e != nil {
		log.Printf("Error: %v", e)
		return nil, false
	}
	log.Println(sid)
	log.Println(sessionCacheByID)
	if s, ok := sessionCacheByID[sid]; ok {
		log.Println(`+***************+`)
		return s, true
	}
	return nil, false
}

func GetSessionByName(name string) (Session, bool) {
	if s, ok := sessionCacheByUserName[name]; ok {
		return *s, true
	}
	return Session{}, false
}

func AddSession(s *Session) {
	if _, ok := sessionCacheByID[s.id]; ok {
		log.Printf("has been exists in cache: %v\n", s)
	} else {
		sessionCacheByID[s.id] = s
		sessionCacheByUserName[s.name] = s
		log.Printf("add new session to cache: %v\n", s)
	}
}

func DeleteSession(uuid guuid.UUID, name string) {
	if _, ok := sessionCacheByID[uuid]; ok {
		delete(sessionCacheByID, uuid)
		log.Printf("has been deleted session from cache_ID: %v\n", uuid)
	}

	if _, ok := sessionCacheByUserName[name]; ok {
		delete(sessionCacheByUserName, name)
		log.Printf("has been deleted session from cache_Name: %v\n", name)
	}
}

func ReleaseSessionCache(now int64) {
	if now-lastReleaseTime < 1000 {
		return
	}

	log.Printf("Start finding invalid session. %d\n", len(sessionCacheByID))

	invalidKeys := []func(){}
	for _, session := range sessionCacheByID {
		if utils.NowMilliseconds()-session.lastHeartBeatMillisecond > 30000 {
			uuid, name := session.Close("Your connection will be forcibly closed later")

			// 放置 goroutine queue
			invalidKeys = append(invalidKeys, func() { DeleteSession(uuid, name) })
			// go DeleteSession(conn, uuid, name)
		}
	}

	log.Printf("Invalid sesisons count %d", len(invalidKeys))
	for _, fun := range invalidKeys {
		fun()
	}
}
