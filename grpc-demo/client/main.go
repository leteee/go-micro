package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-demo/greeter/greeter"
	"log"
	"time"
)

const (
	defaultName = "world"
)

var (
	//addr = flag.String("addr", "localhost:8000", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {

	//初始化consul配置
	consulConfig := api.DefaultConfig()

	//获取consul操作对象
	registry, _ := api.NewClient(consulConfig)

	serviceEntry, _, _ := registry.Health().Service("HelloService", "testHello", false, &api.QueryOptions{})
	addr := fmt.Sprintf("%s:%v", serviceEntry[0].Service.Address, serviceEntry[0].Service.Port)
	fmt.Println("consul获取的地址-", addr)
	flag.Parse()
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
