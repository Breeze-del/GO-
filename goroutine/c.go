package main

import (
	"io"
	"log"
	"net"
	"os"
)

// 客户端
func main() {
	// tcp 连接接口
	coon, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer coon.Close()
	go mustCopy(os.Stdout, coon)
	mustCopy(coon, os.Stdin)
}

// 将coon【服务器】 返回得信息copy输出到os.stdout上
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
