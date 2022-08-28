package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc/credentials"
	"log"
	"time"

	"hello_client/pb"

	"google.golang.org/grpc"
)

// grpc 客户端
// 调用server端的 SayHello 方法

var name = flag.String("name", "七米", "通过-name告诉server你是谁")

func main() {
	flag.Parse() // 解析命令行参数

	// 连接server
	// 加载证书
	creds, err := credentials.NewClientTLSFromFile("certs/server.crt", "liwenzhou.com")
	if err != nil {
		fmt.Printf("credentials.NewClientTLSFromFile failed, err:%v\n", err)
		return
	}
	conn, err := grpc.Dial("127.0.0.1:8972",
		//grpc.WithTransportCredentials(insecure.NewCredentials()), // 不安全的连接
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		log.Fatalf("grpc.Dial failed,err:%v", err)
		return
	}
	defer conn.Close()
	// 创建客户端
	c := pb.NewGreeterClient(conn) // 使用生成的Go代码
	// 调用RPC方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		fmt.Printf("c.SayHello failed, err:%v\n", err)
		return
	}
	// 拿到了RPC响应
	fmt.Printf("resp:%v\n", resp.GetReply())
}
