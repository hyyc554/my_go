package main

import (
	"fmt"
	"net"
	"time"
)

func handle(conn net.Conn) {
	defer conn.Close()
	fmt.Println("connect:", conn.RemoteAddr())
	for {
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		if n == 0 {
			fmt.Printf("%s has dis connect", conn.RemoteAddr())
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Receive %d byte data : %s \n", n, string(data[:n]))
		//程序睡眠1ns，模拟网络时延
		time.Sleep(time.Nanosecond)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8899")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println(err)
			return
		}
		go handle(conn)
	}

}
