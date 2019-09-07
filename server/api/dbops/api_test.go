package dbops

import "testing"

var tempvid string

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")

}
func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}
func TestUserWorkFlow(t *testing.T) {
	t.Run("add", testAddUser)
	t.Run("get", testGetUser)
	t.Run("del", testDelUser)
	t.Run("reget", testReGetUser)
}

func TestVideoWorkFlow(t *testing.T) {
	t.Run("prepareUser", testAddUser)
	t.Run("addvideo", testAddVideo)
	t.Run("getVideo", testGetVideo)
	t.Run("delvideo", testDelVideo)
	t.Run("regetvidoe", testReGetVideo)
}
func testAddUser(t *testing.T) {
	err := AddUserCredential("anjun", "123")
	if err != nil {
		t.Errorf("error of add user:%v", err)
	}
}
func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("anjun")
	if pwd != "123" || err != nil {
		t.Errorf("error of get user:%v,%v", pwd, err)
	}
}
func testDelUser(t *testing.T) {
	err := DeleteUser("anjun", "123")
	if err != nil {
		t.Errorf("error of add user:%v", err)
	}
}
func testReGetUser(t *testing.T) {
	pwd, err := GetUserCredential("anjun")
	if err != nil {
		t.Errorf("error of get user:%v", pwd)
	}
	if pwd != "" {
		t.Errorf("deleting user test failed:%v", pwd)
	}
}
func testAddVideo(t *testing.T) {
	vi, err := addVideo(1, "my-video")
	if err != nil {
		t.Errorf("error of add video:%v", err)
	}
	tempvid = vi.Id
}
func testGetVideo(t *testing.T) {
	vi, err := getVideoInfo(tempvid)
	if err != nil {
		t.Errorf("error of get video:%v", err)
	}
	tempvid = vi.Id
}
func testDelVideo(t *testing.T) {
	err := DeleteVideoInfo("1")
	if err != nil {
		t.Errorf("error of del video:%v", err)
	}

}
func testReGetVideo(t *testing.T) {
	vi, err := getVideoInfo(tempvid)
	if err != nil {
		t.Errorf("error of reget video:%v", err)
	}
	if vi.Name != "my-video" {
		t.Errorf("del error %v", err)
	}
}
