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

	/*
		// 1. 连接consul
		cc, err := api.NewClient(api.DefaultConfig())
		if err != nil {
			fmt.Printf("api.NewClient failed, err:%v\n", err)
			return
		}
		// 2. 根据服务名称查询服务实例
		// cc.Agent().Services()  // 列出所有的
		serviceMap, err := cc.Agent().ServicesWithFilter("Service==`hello`") // 查询服务名称是hello的所有服务节点
		if err != nil {
			fmt.Printf("query `hello` service failed, err:%v\n", err)
			return
		}
		var addr string
		for k, v := range serviceMap {
			fmt.Printf("%s:%#v\n", k, v)
			addr = fmt.Sprintf("%s:%d", v.Address, v.Port) // 取第一个机器的address和port
			continue
		}

		// map[string]*api.AgentService
		// 从consul返回的数据中选一个服务实例（机器）

		// 3. 与consul返回的服务实例建立连接
		// 连接server
		conn, err := grpc.Dial(addr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)

	*/

	conn, err := grpc.Dial(
		"consul://localhost:8500/hello", // grpc中使用consul名称解析器，
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
