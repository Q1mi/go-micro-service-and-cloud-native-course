package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"addsrv/pb"
)

// go-kit addService demo 3

var (
	httpAddr = flag.Int("http-addr", 8080, "HTTP端口")
	gRPCAddr = flag.Int("grpc-addr", 8972, "gRPC端口")
)

func main() {
	// 前置资源初始化

	srv := NewService()

	var g errgroup.Group

	// HTTP
	g.Go(func() error {
		httpListener, err := net.Listen("tcp", fmt.Sprintf(":%d", *httpAddr))
		if err != nil {
			fmt.Printf("net.Listen %d faield, err:%v\n", *httpAddr, err)
			return err
		}
		defer httpListener.Close()
		httpHandler := NewHTTPServer(srv)

		return http.Serve(httpListener, httpHandler)
	})
	// gRPC
	g.Go(func() error {
		grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", *gRPCAddr))
		if err != nil {
			fmt.Printf("net.Listen %d faield, err:%v\n", *gRPCAddr, err)
			return err
		}
		defer grpcListener.Close()

		s := grpc.NewServer() // gRPC Server
		pb.RegisterAddServer(s, NewGRPCServer(srv))
		return s.Serve(grpcListener)
	})

	// wait
	if err := g.Wait(); err != nil {
		fmt.Printf("server exit with error:%v\n", err)
	}
}
