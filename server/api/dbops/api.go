package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jsxz/go-video/server/api/defs"
	"github.com/jsxz/go-video/server/api/utils"
	"log"
	"time"
)

func AddUserCredential(loginName, pwd string) error {
	stmt, err := dbConn.Prepare("insert into users (login_name,pwd) values(?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}
func GetUserCredential(loginName string) (string, error) {
	stmt, err := dbConn.Prepare("select pwd from users where login_name=?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	err = stmt.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmt.Close()
	return pwd, nil
}
func DeleteUser(loginName, pwd string) error {
	stmt, err := dbConn.Prepare("delete from users where login_name=? and pwd=?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	_, err = stmt.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}
func addVideo(aid int, name string) (*defs.VideoInfo, error) {
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05") //M D y,HH:MM:SS
	stmt, err := dbConn.Prepare(`insert into video_info(id,author_id,name,
display_ctime) values (?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	defer stmt.Close()
	return res, nil
}
func getVideoInfo(vid string) (*defs.VideoInfo, error) {

	stmt, err := dbConn.Prepare(`select author_id,name,display_ctime from video_info where id=?`)
	if err != nil {
		return nil, err
	}
	var (
		aid   int
		ctime string
		name  string
	)
	err = stmt.QueryRow(vid).Scan(&aid, &name, &ctime)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	defer stmt.Close()
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}

	return res, nil
}
func DeleteVideoInfo(vid string) error {
	stmt, err := dbConn.Prepare("delete from video_info where id=?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	_, err = stmt.Exec(vid)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}
