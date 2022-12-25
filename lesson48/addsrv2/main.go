package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/go-kit/log"
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
	// 初始化带logger的service
	logger := log.NewJSONLogger(os.Stdout)
	srv = NewLogMiddleware(logger, srv)

	var g errgroup.Group

	// HTTP
	g.Go(func() error {
		httpListener, err := net.Listen("tcp", fmt.Sprintf(":%d", *httpAddr))
		if err != nil {
			fmt.Printf("net.Listen %d faield, err:%v\n", *httpAddr, err)
			return err
		}
		defer httpListener.Close()
		// 初始化go-kit logger
		logger := log.NewLogfmtLogger(os.Stderr)
		httpHandler := NewHTTPServer(srv, logger)

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
