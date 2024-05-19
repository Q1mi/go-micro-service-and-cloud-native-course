package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	_ "github.com/mbobakov/grpc-consul-resolver" // 匿名导入
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"hello_client/pb"
)

// grpc 客户端
// 调用server端的 SayHello 方法

var name = flag.String("name", "七米", "通过-name告诉server你是谁")

func main() {
	flag.Parse() // 解析命令行参数

	conn, err := grpc.Dial(
		"consul://localhost:8500/hello?healthy=true", // grpc中使用consul名称解析器，
		// 指定负载均衡策略，这里使用的是gRPC自带的round_robin
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("grpc.Dial failed,err:%v", err)
		return
	}
	defer conn.Close()
	// 创建客户端
	c := pb.NewGreeterClient(conn) // 使用生成的Go代码
	// 4. 发起RPC调用
	// 调用RPC方法
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
		if err != nil {
			fmt.Printf("c.SayHello failed, err:%v\n", err)
			return
		}

		// 拿到了RPC响应
		fmt.Printf("resp:%v\n", resp.GetReply())
		time.Sleep(time.Millisecond * 500)
	}
}

// 总共4个subConn(子链接)
// next=3， 3+1%4=0
// next=0,  0+1%4=1
// next=1,  1+1%4=2
// next=2,  2+1%4=3
// next=3,  3+1%4=0
