package main

import (
	"log"
	"myapp1/rpcdemo"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 服务端提供服务
func main() {
	rpc.Register(rpcdemo.DemoService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept error: %s", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
