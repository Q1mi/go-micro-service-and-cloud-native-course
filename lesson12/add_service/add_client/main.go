package main

import (
	"context"
	"flag"
	"log"
	"time"

	"add_client/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// add rpc client

var (
	x = flag.Int64("x", 10, "x的值")
	y = flag.Int64("y", 20, "y的值")
)

func main() {
	flag.Parse() // 从命令行解析x和y的值

	// 连接rpc server
	conn, err := grpc.Dial("127.0.0.1:8973", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial failed, err:%v", err)
		return
	}
	defer conn.Close()
	// 创建rpc client端
	client := proto.NewCalcServiceClient(conn)
	// 发起RPC调用
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.Add(ctx, &proto.AddRequest{X: int32(*x), Y: int32(*y)})
	if err != nil {
		log.Fatalf("client.Add failed, err:%v", err)
		return
	}
	// 打印结果
	log.Printf("result:%v\n", resp.GetResult())
}
