package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")
	defer conn.Close()
	checkError(err)

	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			checkError(err)
			fmt.Println(string(buf[:n]))
		}
	}()

	var msg string
	for {
		fmt.Scanln(&msg)

		if msg == "exit" {
			os.Exit(0)
		}

		// 发送数据
		_, err = conn.Write([]byte(msg))
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
