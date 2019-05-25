package main

import (
	"fmt"
	"go_study/day13/demo/common"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Connect :", conn.RemoteAddr())

	fd := common.SocketUtil{conn}
	for {
		data, err := fd.ReadPkg() //读取数据
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(data))
	}

}
func main() {
	listener, err := net.Listen("tcp", ":8899")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Start listen localhost:8899")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handle(conn)
	}
}
