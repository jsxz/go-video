package session

import (
	"github.com/jsxz/go-video/server/api/dbops"
	"github.com/jsxz/go-video/server/api/defs"
	"github.com/jsxz/go-video/server/api/utils"
	"sync"
	"time"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}
func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteVideoInfo(sid)
}
func LoadSessionFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}
func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := nowInMilli()
	ttl := ct + 30*60*1000
	ss := &defs.SimpleSession{UserName: un, TTL: ttl}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, un)
	return id
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}
func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := nowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).UserName, false
	}
	return "", true
}
