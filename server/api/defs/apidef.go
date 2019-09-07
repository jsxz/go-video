package defs

//requests
type UserCredential struct {
	UserName string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

//respose
type SignedUp struct {
	Success   bool   `json:"success"`
	SeesionId string `json:"seesion_id"`
}

//Data model
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}
type Comment struct {
	Id      string
	VideoId string
	Author  string
	Content string
}
type SimpleSession struct {
	UserName string
	TTL      int64
}
