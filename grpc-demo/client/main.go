package main

import (
	"client/proto/greet"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	grpcClient, err := grpc.Dial("127.0.0.1:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	//注册客户端
	client := greet.NewGreetClient(grpcClient)

	res, err := client.SayHello(context.Background(), &greet.HelloReq{
		Name: "李华",
	})

	fmt.Println("%#v", res)
}
