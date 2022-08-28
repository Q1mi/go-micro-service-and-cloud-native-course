package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"hello_server/proto"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

// grpc server

type server struct {
	proto.UnimplementedGreeterServer
}

// SayHello 是我们需要实现的方法
// 这个方法是我们对外提供的服务
func (s *server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {
	// 利用defer 在发送完响应数据后，发送trailer
	defer func() {
		trailer := metadata.Pairs(
			"timestamp", strconv.Itoa(int(time.Now().Unix())),
		)
		grpc.SetTrailer(ctx, trailer)
	}()
	// 在执行业务逻辑之前要check metadata中是否包含token
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Printf("md:%#v ok:%#v\n", md, ok)
	if !ok { // 没有元数据我拒接
		return nil, status.Error(codes.Unauthenticated, "无效请求")
	}
	vl := md.Get("token")
	if len(vl) < 1 || vl[0] != "app-test-q1mi" {
		return nil, status.Error(codes.Unauthenticated, "无效token")
	}
	//if vl, ok := md["token"]; ok {
	//	if len(vl) > 0 && vl[0] == "app-test-q1mi" {
	//		// 有效的请求
	//	}
	//}
	reply := "hello " + in.GetName()
	// 发送数据前发送header
	header := metadata.New(map[string]string{
		"location": "Beijing",
	})
	grpc.SendHeader(ctx, header)
	return &proto.HelloResponse{Reply: reply}, nil
}

func main() {
	// 启动服务
	l, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen, err:%v\n", err)
		return
	}
	s := grpc.NewServer() // 创建grpc服务
	// 注册服务
	proto.RegisterGreeterServer(s, &server{})
	// 启动服务
	err = s.Serve(l)
	if err != nil {
		fmt.Printf("failed to serve,err:%v\n", err)
		return
	}
}
