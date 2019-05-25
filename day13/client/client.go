package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8899") //连接服务端
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connect to localhost:8899 success")
	defer conn.Close()

	for i := 0; i < 100; i++ {
		data := fmt.Sprintf("{\"index\":%d,\"name\":\"maqian\", \"age\":21, \"company\":\"intely\"}", i+1)
		n, err := conn.Write([]byte(data))
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Send %d byte data : %s", n, data)

	}
}
