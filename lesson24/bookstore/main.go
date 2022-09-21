package main

import (
	"bookstore/pb"
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Bookstore

func main() {
	// 连接数据库
	db, err := NewDB("bookstore.db")
	if err != nil {
		fmt.Printf("connect to db failed,err:%v\n", err)
		return
	}

	// 创建server
	srv := server{
		bs: &bookstore{db: db},
	}

	// 启动gRPC服务
	l, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen, err:%v\n", err)
		return
	}
	s := grpc.NewServer()
	// 注册服务
	pb.RegisterBookstoreServer(s, &srv)
	go func() {
		fmt.Println(s.Serve(l))
	}()

	// grpc-Gateway
	conn, err := grpc.DialContext(
		context.Background(),
		"127.0.0.1:8972",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Printf("grpc conn failed, err:%v\n", err)
		return
	}

	gwmux := runtime.NewServeMux()
	pb.RegisterBookstoreHandler(context.Background(), gwmux, conn)

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}
	fmt.Println("grpc-Gateway serve on :8090...")
	gwServer.ListenAndServe()
}
