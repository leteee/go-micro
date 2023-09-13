package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	pb "grpc-demo/greeter/greeter"
	"log"
	"net"
)

var (
	port = flag.Int("port", 8000, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	//初始化consul配置
	consulConfig := api.DefaultConfig()

	//获取consul操作对象
	registry, _ := api.NewClient(consulConfig)

	//注册服务,服务的常规配置
	registerService := api.AgentServiceRegistration{
		ID:      "1",
		Tags:    []string{"testHello"},
		Name:    "HelloService",
		Port:    *port,
		Address: "127.0.01",
		Check: &api.AgentServiceCheck{
			TCP:      fmt.Sprintf("127.0.0.1:%d", *port),
			Timeout:  "5s",
			Interval: "5s",
		},
	}
	//注册服务到consul上
	registry.Agent().ServiceRegister(&registerService)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
