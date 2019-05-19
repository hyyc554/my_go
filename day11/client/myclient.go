package main

import (
	"fmt"
	"net"
)

func main() {
	tcp_addr, err := net.ResolveTCPAddr("tcp4", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	tcp_conn, err := net.DialTCP("tcp4", nil, tcp_addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	send_data := "helloworld"
	n, err := tcp_conn.Write([]byte(send_data))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Send %d byte data success: %s", n, send_data)
	recvData := make([]byte, 2048)
	n, err = tcp_conn.Read(recvData) //读取数据
	if err != nil {
		fmt.Println(err)
		return
	}
	recvStr := string(recvData[:n])
	fmt.Printf("Response data: %s", recvStr)
}
