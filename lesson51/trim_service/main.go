package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"trim_service/pb"

	apiconsul "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

const serviceName = "trim_service"

var (
	port       = flag.Int("port", 8975, "service port")
	consulAddr = flag.String("consul", "localhost:8500", "consul address")
)

// trim service

type server struct {
	pb.UnimplementedTrimServer
}

// TrimSpace 去除字符串参数中的空格
func (s *server) TrimSpace(_ context.Context, req *pb.TrimRequest) (*pb.TrimResponse, error) {
	ov := req.GetS()
	v := strings.ReplaceAll(ov, " ", "")
	fmt.Printf("ov:%s v:%v\n", ov, v)
	return &pb.TrimResponse{S: v}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterTrimServer(s, &server{})

	// 服务注册
	cc, err := NewConsulClient(*consulAddr)
	if err != nil {
		fmt.Printf("failed to NewConsulClient: %v", err)
		return
	}
	ipInfo, err := getOutboundIP()
	if err != nil {
		fmt.Printf("getOutboundIP failed, err:%v\n", err)
		return
	}
	if err := cc.RegisterService(serviceName, ipInfo.String(), *port); err != nil {
		fmt.Printf("regToConsul failed, err:%v\n", err)
		return
	}
	go func() {
		if err := s.Serve(lis); err != nil {
			fmt.Printf("failed to serve: %v", err)
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	// 退出时注销服务
	cc.Deregister(fmt.Sprintf("%s-%s-%d", serviceName, ipInfo.String(), *port))
}

// consul reg&de
type consulClient struct {
	client *apiconsul.Client
}

// NewConsulClient 新建consulClient
func NewConsulClient(consulAddr string) (*consulClient, error) {
	cfg := apiconsul.DefaultConfig()
	cfg.Address = consulAddr
	client, err := apiconsul.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &consulClient{client}, nil
}

// RegisterService 服务注册
func (c *consulClient) RegisterService(serviceName, ip string, port int) error {
	srv := &apiconsul.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", serviceName, ip, port), // 服务唯一ID
		Name:    serviceName,                                    // 服务名称
		Tags:    []string{"q1mi", "trim"},                       // 为服务打标签
		Address: ip,
		Port:    port,
	}
	return c.client.Agent().ServiceRegister(srv)
}

// Deregister 注销服务
func (c *consulClient) Deregister(serviceID string) error {
	return c.client.Agent().ServiceDeregister(serviceID)
}

// getOutboundIP 获取本机的出口IP
func getOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}
