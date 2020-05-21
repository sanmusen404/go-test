package main

import (
	"fmt"
	"net"
	"os"
)

const MAXINUM_CLIENT = 10

func main() {
	listener, err := net.Listen("tcp", ":8888")
	checkError(err)
	defer listener.Close()

	fmt.Println("服务器已启动，等待客户端建立连接...")

	count := make(chan int, MAXINUM_CLIENT)

	for {
		count <- 1
		go func() {
			conn, err := listener.Accept()
			checkError(err)
			defer conn.Close()
			ip := conn.RemoteAddr().String()
			fmt.Println("系统消息：" + ip + "加入聊天室")
			for {
				buf := make([]byte, 1024) // 创建1024大小的缓冲区，用于read
				n, err := conn.Read(buf)
				if err != nil {
					fmt.Println("系统消息：" + ip + "退出了聊天室")
					break
				}
				fmt.Println(ip+":", string(buf[:n])) // 读多少，打印多少。
			}

			<-count
		}()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
