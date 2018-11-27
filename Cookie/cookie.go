package main

import (
	"io"
	"net/http"
	"strings"
)

// 不用框架，Go原生操作cookie方法
func main() {
	http.HandleFunc("/", Cookie2)
	http.ListenAndServe(":8080", nil)
}

func Cookie(w http.ResponseWriter, r *http.Request) {
	// 申请一个cookie结构体
	ck := &http.Cookie{
		Name:   "myCookie",
		Value:  "hello",
		Path:   "/",
		Domain: "localhost",
		MaxAge: 120,
	}
	// 在response中设置了cookie
	http.SetCookie(w, ck)
	// 从request中获取cookie
	ck2, err := r.Cookie("myCookie")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, ck2.Value)
}

func Cookie2(w http.ResponseWriter, r *http.Request) {
	// 申请一个cookie结构体
	// cookie的value中不能存在' '空格
	ck := &http.Cookie{
		Name:   "myCookie",
		Value:  "hello world",
		Path:   "/",
		Domain: "localhost",
		MaxAge: 120,
	}
	w.Header().Set("Set-Cookie", strings.Replace(ck.String(), " ", "%20", -1))
	ck2, err := r.Cookie("myCookie")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, ck2.Value)
}
