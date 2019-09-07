package main

import (
	"io"
	"net/http"
)

func Print1to20() int {
	res := 0
	for i := 0; i <= 20; i++ {
		res += i
	}
	return res
}
func main() {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":8888", nil)
}

func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>hello world</h1>")
}
