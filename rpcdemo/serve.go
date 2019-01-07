package main

import (
	"github.com/pkg/errors"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type DemoService struct {
}

type Args struct {
	A, B int
}

// 一个服务
func (d *DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	*result = float64(args.A) / float64(args.B)
	return nil
}

func main() {
	rpc.Register(&DemoService{})
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
