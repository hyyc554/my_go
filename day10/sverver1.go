package main

import (
	"fmt"
	"net"
)

//func handleConn(c net.Conn) {
//	defer c.Close()
//	for {
//		// read from the connection58/*
//		// ... ...
//		// write to the connection
//		//... ...
//	}
//}
func handle(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 100)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("read data size %d msg:%s\n", n, string(buf[0:n]))
		//msg := []byte("hello,world\n")
		conn.Write(buf[0:n]) //发送数据

	}

}

func main() {
	fmt.Println("start server....")
	listen, err := net.Listen("tcp", "0.0.0.0:3000") //创建监听
	if err != nil {
		fmt.Println("listen failed! msg :", err)
		return
	}
	for {
		conn, errs := listen.Accept() //接受客户端连接
		if errs != nil {
			fmt.Println("accept failed")
			continue
		}
		go handle(conn) //处理连接
	}
}
