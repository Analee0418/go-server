package model

import (
	"log"

	guuid "github.com/google/uuid"
)

var sessionCache = map[guuid.UUID]*Session{}

var sessionCacheByUserName = map[string]*Session{}

func GetSessionByName(name string) (Session, bool) {
	if s, ok := sessionCacheByUserName[name]; ok {
		return *s, true
	}
	return Session{}, false
}

func AddSession(s *Session) {
	if _, ok := sessionCache[s.id]; ok {
		log.Println("has been exists in cache: %v", s)
	} else {
		sessionCache[s.id] = s
		log.Println("add new session to cache: %v", s)
	}
}

func DeleteSession(s *Session) {
	if _, ok := sessionCache[s.id]; ok {
		delete(sessionCache, s.id)
		log.Println("has been deleted session from cache: %v", s)
	}
}
