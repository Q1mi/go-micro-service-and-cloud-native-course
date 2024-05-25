package main

import (
	"context"
	"fmt"
	"net"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	"addsrv/pb"
)

// go-kit addService demo 2

// 1. service
// 1.1 业务逻辑抽象为接口
type AddService interface {
	Sum(ctx context.Context, a, b int) (int, error)
}
// 1.2 实现接口
type addService struct{}
func (addService) Sum(_ context.Context, a, b int) (int, error) {
	return a + b, nil
}
// 1.3 请求和响应
type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}
type SumResponse struct {
	V   int    `json:"v"`
	Err string `json:"err,omitempty"`
}

// 2. endpoint
// 借助 适配器 将 方法 -> endpoint
func makeSumEndpoint(srv AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SumRequest)
		v, err := srv.Sum(ctx, req.A, req.B) // 方法调用
		if err != nil {
			return SumResponse{V: v, Err: err.Error()}, nil
		}
		return SumResponse{V: v}, nil
	}
}

// 3. transport
type grpcServer struct {
	pb.UnimplementedAddServer
	sum    grpctransport.Handler
}
func NewGRPCServer(svc AddService) pb.AddServer {// NewGRPCServer 构造函数
	return &grpcServer{
		sum: grpctransport.NewServer(
			makeSumEndpoint(svc), // endpoint
			decodeGRPCSumRequest,
			encodeGRPCSumResponse,
		),
	}
}
// decodeGRPCSumRequest 将Sum方法的gRPC请求参数转为内部的SumRequest
func decodeGRPCSumRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.SumRequest)
	return SumRequest{A: int(req.A), B: int(req.B)}, nil
}
// encodeGRPCSumResponse 封装Sum的gRPC响应
func encodeGRPCSumResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(SumResponse)
	return &pb.SumResponse{V: int64(resp.V), Err: resp.Err}, nil
}
func (s grpcServer) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	_, resp, err := s.sum.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.SumResponse), nil
}

func main() {
	srv := addService{}
	gs := NewGRPCServer(srv)

	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Printf("net.Listen faield, err:%v\n", err)
		return
	}

	s := grpc.NewServer() // gRPC Server
	pb.RegisterAddServer(s, gs)

	// 启动服务
	fmt.Println(s.Serve(listener))
}
