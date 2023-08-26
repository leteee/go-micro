package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	//1.用rpc.Dial和rpc建立连接
	conn, err := rpc.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
	}
	//关闭连接
	defer conn.Close()

	//调用函数
	var reply string
	err = conn.Call("hello.SayHello", "客户端", &reply)
	if err != nil {
		fmt.Println(err)
	}
	//获取服务端返回的数据
	fmt.Println(reply)
}
