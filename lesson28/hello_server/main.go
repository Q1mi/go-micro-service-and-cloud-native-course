package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	"hello_server/pb"

	"google.golang.org/grpc"
)

// grpc server

var port = flag.Int("port", 8972, "服务端口")

type server struct {
	pb.UnimplementedGreeterServer
	Addr string
}

// SayHello 是我们需要实现的方法
// 这个方法是我们对外提供的服务
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	reply := fmt.Sprintf("hello %s. [from %s]", in.GetName(), s.Addr)
	return &pb.HelloResponse{Reply: reply}, nil
}

func main() {
	flag.Parse()
	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	// 启动服务
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("failed to listen, err:%v\n", err)
		return
	}

	s := grpc.NewServer() // 创建grpc服务
	// 注册服务
	pb.RegisterGreeterServer(s, &server{Addr: addr})
	// 启动服务
	err = s.Serve(l)
	if err != nil {
		fmt.Printf("failed to serve,err:%v\n", err)
		return
	}
}
