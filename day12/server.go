package main

import (
	"fmt"
	"net"
	"strings"
)

func handle(conn net.Conn) {
	defer conn.Close()
	fmt.Println("connect:", conn.RemoteAddr())
	for {
		data := make([]byte, 2048)
		n, err := conn.Read(data)
		if n == 0 {
			fmt.Printf("%s has dis connect", conn.RemoteAddr())
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Receive data [%s] from [%s]", string(data[:n]), conn.RemoteAddr())
		rspData := strings.ToUpper(string(data[:n]))
		_, err = conn.Write([]byte(rspData))
		if err != nil {
			fmt.Println(err)
			continue
		}
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
