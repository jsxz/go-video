package dbops

import (
	"database/sql"
	"github.com/jsxz/go-video/server/api/defs"
	"log"
	"strconv"
	"sync"
)

func InsertSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmt, err := dbConn.Prepare("insert into sessions(session_id,TTL,login_name)values(?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}
func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmt, err := dbConn.Prepare("select TTL,login_name from sessions where id=?")
	if err != nil {
		return nil, err
	}
	var ttlstr string
	var uname string
	stmt.QueryRow(sid).Scan(&ttlstr, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if ttl, err := strconv.ParseInt(ttlstr, 10, 64); err == nil {
		ss.TTL = ttl
		ss.UserName = uname
	} else {
		return nil, err
	}
	defer stmt.Close()
	return ss, nil
}
func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmt, err := dbConn.Prepare("select TTL,login_name from sessions")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	for rows.Next() {
		var id string
		var ttlstr string
		var login_name string
		if er := rows.Scan(&id, &ttlstr, &login_name); er != nil {
			log.Println("retrive sessions error : %v", er)
			break
		}
		if ttl, err := strconv.ParseInt(ttlstr, 10, 64); err == nil {
			ss := &defs.SimpleSession{UserName: login_name, TTL: ttl}
			m.Store(id, ss)
			log.Printf("session id :%s,ttl: %d ", id, ttl)
		} else {
			return nil, err
		}
	}

	defer stmt.Close()
	return m, nil
}
func DeleteSession(sid string) error {
	stmt, err := dbConn.Prepare("delete from sessions where id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sid)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}
