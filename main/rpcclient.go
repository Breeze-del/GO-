package main

import (
	"fmt"
	"myapp1/rpcdemo"
	"net"
	"net/rpc/jsonrpc"
)

// 模拟客户端连接
func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)
	var result float64
	err1 := client.Call("DemoService.Div", rpcdemo.Args{10, 3},
		&result)
	fmt.Println(result, err1)
	err2 := client.Call("DemoService.Div", rpcdemo.Args{10, 0},
		&result)
	fmt.Println(result, err2)
}
