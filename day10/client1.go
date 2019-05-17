package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println("err dialing:", err.Error())
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		str, _ := inputReader.ReadString('\n')
		data := strings.Trim(str, "\n")
		if data == "quit" { //输入quit退出
			return
		}
		_, err := conn.Write([]byte(data)) //发送数据
		if err != nil {
			fmt.Println("send data error:", err)
			return
		}
		buf := make([]byte, 512)
		n, err := conn.Read(buf) //读取服务端端数据
		fmt.Println("from server:", string(buf[:n]))
	}
}
