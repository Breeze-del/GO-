package main

import (
	"net/http"
	"net/http/httptest"
)

// HTTP 中间件简单开发
type MyHandler struct {
	mHandler   http.Handler
	signalHost string
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// rec记录正常响应所写入的所有东西
	rec := httptest.NewRecorder()
	// 检验host是否正确
	if req.Host == this.signalHost {
		this.mHandler.ServeHTTP(rec, req)
		// 将rec中header写入response
		for k, v := range rec.HeaderMap {
			w.Header()[k] = v
		}
		// 自定义写入头部或者body
		w.Header().Add("kaitou", "wsnd")
		w.Write(rec.Body.Bytes())
		w.Write([]byte("追加写入"))
		w.WriteHeader(418)
	} else {
		w.Write([]byte("no mathcing"))
	}
}

func serve(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("dddd dddd ddddd"))
}
func main() {
	myhander := &MyHandler{
		mHandler:   http.HandlerFunc(serve),
		signalHost: "localhost:8080",
	}
	serve := http.Server{
		Handler: myhander,
		Addr:    "localhost:8080",
	}
	serve.ListenAndServe()
}
