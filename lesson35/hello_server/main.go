package main

import (
	"context"
	"fmt"
	"net"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"

	"hello_server/pb"

	"google.golang.org/grpc/health"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// grpc server

const serviceName = "hello"

type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello 是我们需要实现的方法
// 这个方法是我们对外提供的服务
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	reply := "hello " + in.GetName()
	return &pb.HelloResponse{Reply: reply}, nil
}

func main() {
	// 启动服务
	l, err := net.Listen("tcp", ":8976")
	if err != nil {
		fmt.Printf("failed to listen, err:%v\n", err)
		return
	}

	s := grpc.NewServer() // 创建grpc服务
	// 注册服务
	pb.RegisterGreeterServer(s, &server{})
	// 给我们的gRPC服务增加了健康检查的处理逻辑
	healthpb.RegisterHealthServer(s, health.NewServer()) // consul 发来健康检查的RPC请求，这个负责返回OK

	// 连接至consul
	cc, err := api.NewClient(api.DefaultConfig()) // 127.0.0.1:8500
	if err != nil {
		fmt.Printf("api.NewClient failed, err:%v\n", err)
		return
	}
	// 获取本机的出口ip
	ipinfo, err := GetOutboundIP()
	if err != nil {
		fmt.Printf("GetOutboundIP failed,err:%v\n", err)
		return
	}
	fmt.Println(ipinfo.String())
	// 将我们的gRPC服务注册到consul
	// 1.定义我们的服务
	// 配置健康检查策略，告诉consul如何进行健康检查
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", ipinfo.String(), 8976), // 外网地址
		Timeout:                        "5s",
		Interval:                       "5s",  // 间隔
		DeregisterCriticalServiceAfter: "10s", // 10秒钟后注销掉不健康的服务节点
	}

	srv := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", serviceName, ipinfo.String(), 8976), // 服务唯一ID
		Name:    serviceName,
		Tags:    []string{"qimi"},
		Address: ipinfo.String(),
		Port:    8976,
		Check:   check,
	}
	// 2.注册服务到consul
	cc.Agent().ServiceRegister(srv)

	// 启动服务
	err = s.Serve(l)
	if err != nil {
		fmt.Printf("failed to serve,err:%v\n", err)
		return
	}
}

// GetOutboundIP 获取本机的出口IP
func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}

// 10.22.33.4:80/health   --> HTTP 200
// 127.0.0.1:8972/health  --> gRPC ok
