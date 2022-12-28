package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"strings"

	"trim_service/pb"
)

var port = flag.Int("port", 8975, "service port")

// trim service

type server struct {
	pb.UnimplementedTrimServer
}

// TrimSpace 去除字符串参数中的空格
func (s *server) TrimSpace(_ context.Context, req *pb.TrimRequest) (*pb.TrimResponse, error) {
	ov := req.GetS()
	v := strings.ReplaceAll(ov, " ", "")
	fmt.Printf("ov:%#v v:%#v\n", ov, v)
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
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
