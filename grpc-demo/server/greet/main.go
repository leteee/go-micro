package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"greet/proto/greet"
	"log"
	"net"
)

type Hello struct {
	greet.UnimplementedGreetServer
}

func (h Hello) SayHello(ctx context.Context, req *greet.HelloReq) (*greet.HelloRes, error) {
	fmt.Println(req)
	return &greet.HelloRes{
		Message: "你好，" + req.Name,
	}, nil
}

func main() {
	// 初始化GRPC对象
	grpcServer := grpc.NewServer()

	//注册服务
	greet.RegisterGreetServer(grpcServer, &Hello{})

	//设置监听
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
	}
	//启动服务
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
