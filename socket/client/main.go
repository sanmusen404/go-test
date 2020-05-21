package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	var msg string
	conn, err := net.Dial("tcp", ":8888")
	defer conn.Close()
	checkError(err)

	for {
		fmt.Print("输入信息:")
		fmt.Scanln(&msg)

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
