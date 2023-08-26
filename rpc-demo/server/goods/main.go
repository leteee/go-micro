package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// Goods 创建远程调用的函数，函数一般是放在接口体中
type Goods struct{}

type AddGoodsReq struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

type AddGoodsRes struct {
	Success bool
	Message string
}

func (g Goods) AddGoods(req AddGoodsReq, res *AddGoodsRes) error {
	fmt.Printf("%#v\n", req)
	res = &AddGoodsRes{
		Success: true,
		Message: "数据增加成功",
	}
	return nil
}

func main() {
	//注册RPC服务
	err := rpc.RegisterName("goods", Goods{})
	if err != nil {
		fmt.Println(err)
	}
	//监听端口
	listener, _ := net.Listen("tcp", ":8000")

	defer listener.Close()

	for {
		//接收连接
		fmt.Println("准备建立连接....")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept() err:", err)
		}
		//绑定服务
		go rpc.ServeConn(conn)
	}
}
