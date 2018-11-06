package main

import (
	"io"
	"log"
	"net"
	"time"
)

// 服务
func main() {
	// 监听 8000接口
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		// 如果有人连接服务器 那么返回连接coon【代表客户端】
		coon, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}
		handleConn(coon)
	}
}
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// 向客户端写时间信息
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
