package model

import (
	"log"
	"net"

	"com.lueey.shop/config"
	avro "com.lueey.shop/protocol"
	"com.lueey.shop/utils"
	guuid "github.com/google/uuid"
)

var lastReleaseTime int64 = utils.NowMilliseconds()

var sessionCacheByUserName = map[string]*Session{}
var sessionCacheByID = map[guuid.UUID]*Session{}
var sessionConn = map[net.Conn]*Session{}

var hostsSession *Session = nil

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
}

func Onlines() int32 {
	return int32(len(sessionCacheByID))
}

func SessionByID(uuid string) (*Session, bool) {
	if len(uuid) == 0 {
		return nil, false
	}
	sid, e := guuid.Parse(uuid)
	if e != nil {
		log.Printf("\033[1;31m[ERROR] \033[0m%v", e)
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

func GetSessionByConn(conn net.Conn) (*Session, bool) {
	if s, ok := sessionConn[conn]; ok {
		return s, true
	}
	return nil, false
}

func AddSession(conn net.Conn, s *Session) {
	if _, ok := sessionCacheByID[s.id]; ok {
		log.Printf("[INFO] has been exists in cache: %v\n", s)
	} else {
		sessionCacheByID[s.id] = s
		sessionCacheByUserName[s.name] = s
		sessionConn[conn] = s
		log.Printf("[INFO] add new session to cache: %v\n", s)
	}
}

func DeleteSession(uuid guuid.UUID, name string, conn net.Conn) {
	defer func() { conn.Close() }()
	if _, ok := sessionCacheByID[uuid]; ok {
		delete(sessionCacheByID, uuid)
		log.Printf("[INFO] has been deleted session from cache_ID: %v\n", uuid)
	}

	if _, ok := sessionCacheByUserName[name]; ok {
		delete(sessionCacheByUserName, name)
		log.Printf("[INFO] has been deleted session from cache_Name: %v\n", name)
	}

	if _, ok := sessionConn[conn]; ok {
		delete(sessionConn, conn)
		log.Printf("[INFO] has been deleted session from cache_Conn: %v\n", conn)
	}

	log.Printf("[INFO] session.length: id.%d, name.%d, conn.%d", len(sessionCacheByID), len(sessionCacheByUserName), len(sessionConn))

	if t := len(sessionCacheByID) + len(sessionCacheByUserName) - len(sessionConn); t != len(sessionCacheByID) && t != len(sessionCacheByUserName) && t != len(sessionConn) {
		log.Printf("\033[1;33m[WARNING] \033[0mSession cache conflict!!! id.%d, name.%d, conn.%d", len(sessionCacheByID), len(sessionCacheByUserName), len(sessionConn))
	}
}

func ReleaseSessionCache(now int64) {
	if now-lastReleaseTime < 1000 {
		return
	}

	type inte struct {
		Sid guuid.UUID
		Sn  string
		Sc  net.Conn
	}

	invalidKeys := []inte{}
	for _, session := range sessionCacheByID {
		if session.Dead() {
			invalidKeys = append(invalidKeys, inte{session.id, session.name, session.conn})
			log.Printf("[DEBUG] add dead session %v, %v, %v", session.id, session.name, session.conn)
		} else {
			if utils.NowMilliseconds()-session.lastHeartBeatMillisecond > func() int64 {
				if config.DEBUG {
					return 60000
				}
				return 10000
			}() {
				invalidKeys = append(invalidKeys, inte{session.id, session.name, session.conn})
				session.Close("Connection without vital signs. No heartbeat detected.")
			}
		}
	}

	if len(invalidKeys) > 0 {
		lang, err := json.MarshalIndent(invalidKeys, "", "   ")
		if err == nil {
			log.Printf("[INFO] invalid session keys: %v", string(lang))
		}
		for _, i := range invalidKeys {
			DeleteSession(i.Sid, i.Sn, i.Sc)
		}

		TCPServerInstance.TCPServerOnUpdateOnlines(Onlines(), false)
	}

	lastReleaseTime = now
}

func OnClientDisconnect(conn net.Conn) {
	defer func() { conn.Close() }()
	s, ok := sessionConn[conn]
	if ok {
		if s.customerInfo != nil {
			s.customerInfo = nil
		}
		s.conn.Close()
		s.SetDead()
		log.Printf("[INFO] Client disconnect after clear session[%s].", s.name)
	}

}

func BroadcastMessage(msg avro.Message) {
	for _, session := range sessionConn {
		if !session.dead {
			session.SendMessage(msg)
		}
	}
}

func OnBroadcastToGlobal() {
	pubsub := utils.GetRDB().Subscribe("global_broadcast")
	defer func() { pubsub.Close() }()

	ch := pubsub.ChannelSize(1)
	for {
		res := <-ch
		log.Printf("[INFO] recived broadcast message: %s", res.Payload)

		var message map[string]string
		if err := json.Unmarshal([]byte(res.Payload), &message); err == nil {
			log.Printf("INFO recived broadcast message: %s", message)
			if key, ok := message["key"]; ok {
				if sec, ok := message["sec"]; ok {
					msg := GenerateMessage(avro.ActionMessage_broadcast)
					msg.Message_broadcast = &avro.Message_broadcastUnion{
						UnionType: avro.Message_broadcastUnionTypeEnumMessageForward,
						MessageForward: &avro.MessageForward{
							Key: &avro.KeyUnion{
								UnionType: avro.KeyUnionTypeEnumString,
								String:    key,
							},
							Sec: &avro.SecUnion{
								UnionType: avro.SecUnionTypeEnumString,
								String:    sec,
							},
						},
					}

					for _, session := range sessionConn {
						if !session.dead {
							session.SendMessage(*msg)
						}
					}
				}
			}
		}
	}
}
