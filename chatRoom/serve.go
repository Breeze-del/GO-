package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// 并发聊天室 -- serve端

// 用户实体
type Client struct {
	C    chan string // 用户发送数据的管道
	Name string      // 用户名字
	Addr string      // 用户网络地址
}

// 保存在线用户
var onlineMap map[string]Client

// 公共广播信息通道
var message chan string

func init() {
	message = make(chan string)
	onlineMap = make(map[string]Client)
}

func main() {
	// 监听接口
	listener, err := net.Listen("tcp", "127.0.0.1:888")
	if err != nil {
		log.Fatal("net listen failed")
	}
	go Manager()
	defer listener.Close()
	// 接受连接，并起一个协程处理用户聊天
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listen accept:", err)
			continue
		}
		// 处理用户的练级
		go HandleConn(conn)
	}
}

// 处理用户的连接
func HandleConn(conn net.Conn) {
	defer conn.Close()
	// 获取用户地址
	clientAddr := conn.RemoteAddr().String()
	// 创建用户实体
	client := Client{
		C:    make(chan string),
		Name: clientAddr,
		Addr: clientAddr,
	}
	// 将用户添加到map里
	onlineMap[clientAddr] = client
	// 新建一个协程，专门为当前客户端发送信息
	go WriteMsgToClient(&client, conn)
	// 广播某个用户在线
	message <- MakeStr(&client, "login")
	// 提示客户端，自己是谁
	client.C <- MakeStr(&client, "I am here")
	// 判断是否主动退出聊天室
	isQuit := make(chan bool)
	// 判断是否超时
	isTimeOut := make(chan bool)

	// 处理用户操作的匿名函数
	go func() {
		buf := make([]byte, 1024*2)
		for {
			// 获取客户端的输入
			n, err := conn.Read(buf)
			if n == 0 { // 用户断开连接
				isQuit <- true
				fmt.Println("conn read", err)
				return
			}
			// 获取用户信息
			userMssg := string(buf[:n-1]) // 去掉换行符 "\n"
			// 处理特殊字符
			if len(userMssg) == 3 && userMssg == "who" {
				// 遍历map给当前客户端发送在线的所有客户端
				conn.Write([]byte("user list:\n"))
				for _, tmp := range onlineMap {
					userMssg = tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(userMssg))
				}
			} else if len(message) >= 8 && userMssg[:6] == "rename" {
				// 更名格式 rename|cnm
				newName := strings.Split(userMssg, "|")[1]
				// 更新用户实体和在线用户map
				client.Name = newName
				onlineMap[clientAddr] = client
				// 给客户端回复完成修改
				client.C <- MakeStr(&client, "rename successfully")
			} else {
				// 转发消息
				message <- MakeStr(&client, userMssg)
			}

			// 数据到达此处 说明不超市
			isTimeOut <- true
		}
	}()
	// 判断客户端聊天状态
	for {
		select {
		case <-isQuit:
			delete(onlineMap, clientAddr)
			message <- MakeStr(&client, "logout")
			return
		case <-isTimeOut:
			// 意思是用户要是 10秒内不发消息，那么就会超时
		case <-time.After(10 * time.Second):
			delete(onlineMap, clientAddr)
			message <- MakeStr(&client, "timeout")
			return
		}
	}
}

// 公共消息处理
func Manager() {
	for {
		msg := <-message
		for _, v := range onlineMap {
			v.C <- msg
		}
	}
}

// 给当前客户端发送信息
func WriteMsgToClient(cli *Client, conn net.Conn) {
	// 给当前客户端发送信息
	for msg := range (*cli).C {
		conn.Write([]byte(msg + "\n"))
	}
}

// 制作格式字符串
func MakeStr(cli *Client, str string) string {
	return "{" + (*cli).Addr + "}" + (*cli).Name + ":" + str
}
