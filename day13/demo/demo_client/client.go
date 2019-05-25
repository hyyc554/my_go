package main

import (
	"fmt"
	"go_study/day13/demo/common"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8899")
	if err != nil {
		fmt.Println(err)
		return
	}
	clntFd := common.SocketUtil{conn}
	for i := 0; i < 100; i++ {
		data := fmt.Sprintf("{\"index\":%d, \"name\":\"maqian\", \"age\":21, \"company\":\"intely\"}", i+1)
		n, err := clntFd.WritePkg([]byte(data))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Send %d byte data : %s", n, data)
	}
}
