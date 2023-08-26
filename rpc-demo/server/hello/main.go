package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// Hello 定义一个远程调用的方法
type Hello struct {
}

// SayHello
// 1.方法只能有两个可序列化的参数，其中第二个参数是指针类型
//
//	req 表示获取客户端传过来的数据
//	res 表示给客户端返回数据
//
// 2.方法要返回一个error类型，同时必须是公开的方法
// 3.req和res的类型比如：channel(通道)、func（函数）均不能进行序列化
func (h Hello) SayHello(req string, res *string) error {
	*res = "你好：" + req
	return nil
}

func main() {
	//注册RPC服务
	err := rpc.RegisterName("hello", Hello{})
	if err != nil {
		fmt.Println(err)
	}
	//监听端口
	listener, _ := net.Listen("tcp", ":8000")

	defer listener.Close()

	for {
		//接收连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept() err:", err)
		}
		//绑定服务
		go rpc.ServeConn(conn)
	}

}
