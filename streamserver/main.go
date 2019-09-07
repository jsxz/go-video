package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if !m.l.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "too many request")
		return
	}
	m.r.ServeHTTP(w, r)
	//视频没播放完也先回了，控制不了
	defer m.l.ReleaseConn()
}
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:vid", streamHandler)
	router.POST("/upload/:vid", uploadHandler)
	return router
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r, 2)
	http.ListenAndServe(":9999", mh)
}
