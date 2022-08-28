package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"code.xxx.com/backend/hello_client/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// grpc 客户端
// 调用server端的 SayHello 方法

var name = flag.String("name", "七米", "通过-name告诉server你是谁")

func main() {
	flag.Parse() // 解析命令行参数

	// 连接server
	conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial failed,err:%v", err)
		return
	}
	defer conn.Close()
	// 创建客户端
	c := pb.NewGreeterClient(conn) // 使用生成的Go代码
	// 调用RPC方法
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	// 普通的RPC调用
	//resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	//if err != nil {
	//	log.Printf("c.SayHello failed, err:%v", err)
	//	return
	//}
	//// 拿到了RPC响应
	//log.Printf("resp:%v", resp.GetReply())
	//callLotsOfReplies(c)

	//callLotsOfGreetings(c)

	runBidiHello(c)
}

func callLotsOfReplies(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 调用服务端流式的RPC
	stream, err := c.LotsOfReplies(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Println(err)
		return
	}
	// 依次从流式响应中读取返回的响应数据
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("stream.Recv failed, err:%v\n", err)
			return
		}
		log.Printf("recv :%v\n", res.GetReply())
	}

}

func callLotsOfGreetings(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 客户端要流式的发送请求消息
	stream, err := c.LotsOfGreetings(ctx)
	if err != nil {
		log.Printf("c.LotsOfGreetings(ctx) failed, err:%v\n", err)
		return
	}
	names := []string{"张三", "李四", "王二麻子"}
	for _, name := range names {
		stream.Send(&pb.HelloRequest{Name: name})
	}
	// 流式发送结束之后要关闭流
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("stream.CloseAndRecv() failed, err:%v\n", err)
		return
	}
	log.Printf("res:%v\n", res.GetReply())
}

func runBidiHello(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	// 双向流模式
	stream, err := c.BidiHello(ctx)
	if err != nil {
		log.Fatalf("c.BidiHello failed, err: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			// 接收服务端返回的响应
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("c.BidiHello stream.Recv() failed, err: %v", err)
			}
			fmt.Printf("AI：%s\n", in.GetReply())
		}
	}()
	// 从标准输入获取用户输入
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	for {
		cmd, _ := reader.ReadString('\n') // 读到换行
		cmd = strings.TrimSpace(cmd)
		if len(cmd) == 0 {
			continue
		}
		if strings.ToUpper(cmd) == "QUIT" {
			break
		}
		// 将获取到的数据发送至服务端
		if err := stream.Send(&pb.HelloRequest{Name: cmd}); err != nil {
			log.Fatalf("c.BidiHello stream.Send(%v) failed: %v", cmd, err)
		}
	}
	stream.CloseSend()
	<-waitc
}
