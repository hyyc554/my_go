package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	tcp_linstener, err := net.ListenTCP("tcp4", tcpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("start listen:[%s]", tcpAddr)
	tcp_conn, err := tcp_linstener.AcceptTCP()
	if err != nil {
		fmt.Println(err)
	}
	defer tcp_conn.Close()
	data := make([]byte, 2048)
	n, err := tcp_conn.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	recv_str := string(data[:n])
	fmt.Println("RECV：", recv_str)
	tcp_conn.Write([]byte(strings.ToUpper(recv_str)))
}
