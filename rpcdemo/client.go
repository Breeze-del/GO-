package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

type args struct {
	A, B int
}

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)
	var result float64
	err1 := client.Call("DemoService.Div", args{10, 3},
		&result)
	fmt.Println(result, err1)
}
