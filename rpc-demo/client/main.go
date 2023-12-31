package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

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

func main() {
	//1.用rpc.Dial和rpc建立连接
	//nc -l 169.254.42.250 8000
	//conn, err := rpc.Dial("tcp", "169.254.42.250:8000")
	conn, err := net.Dial("tcp", "169.254.42.250:8000")
	if err != nil {
		fmt.Println(err)
	}
	//关闭连接
	defer conn.Close()

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	//调用函数
	reply := &AddGoodsRes{}
	err = client.Call("goods.AddGoods", AddGoodsReq{
		Id:      1,
		Title:   "追忆似水年华",
		Price:   20,
		Content: "详情",
	}, &reply)
	if err != nil {
		fmt.Println(err)
	}
	//获取服务端返回的数据
	fmt.Printf("%#v\n", reply)
}
