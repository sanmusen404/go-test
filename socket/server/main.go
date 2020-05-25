package main

import (
	"fmt"
	"net"
	"os"
)

const maximumConnection = 10

func main() {
	clients := make(map[string]net.Conn)

	listener, err := net.Listen("tcp", ":8888")
	checkError(err)
	defer listener.Close()

	fmt.Println("\033[32m 服务器已启动，等待客户端建立连接...\033[0m")

	countCh := make(chan int, maximumConnection)
	msgCh := make(chan string)

	//监控消息通道,有新消息时发送到客户端
	go func() {
		for {
			message := <-msgCh
			for _, conn := range clients {
				conn.Write([]byte(message))
			}
		}
	}()

	for {
		countCh <- 1
		go func() {
			conn, err := listener.Accept()
			checkError(err)
			defer conn.Close()

			ip := conn.RemoteAddr().String()
			clients[ip] = conn
			defer delete(clients, ip)
			msgCh <- "\033[31m系统消息：" + ip + "加入聊天室\033[0m"
			for {
				buf := make([]byte, 1024) // 创建1024大小的缓冲区，用于read
				n, err := conn.Read(buf)
				if err != nil {
					msgCh <- "\033[31m系统消息：" + ip + "退出了聊天室\033[0m"
					break
				}
				msgCh <- "@" + ip + ":" + string(buf[:n]) // 读取内容放入消息通道
			}

			<-countCh
		}()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
