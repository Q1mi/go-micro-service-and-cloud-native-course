package main

import (
	"addsrv/pb"
	"context"
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

// gRPC test

// 编写一个gRPC客户端，测试我们的gRPC Server是否正常

// 使用bufconn构建测试链接，避免使用实际端口号启动服务

const bufSize = 1024 * 1024

var bufListener *bufconn.Listener

func init() {
	bufListener = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	gs := NewGRPCServer(addService{})
	pb.RegisterAddServer(s, gs)
	go func() {
		if err := s.Serve(bufListener); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return bufListener.Dial()
}

// 测试代码
func TestSum(t *testing.T) {
	// 1.建立链接
	conn, err := grpc.DialContext(
		context.Background(),
		"bufnet",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(bufDialer),
	)
	if err != nil {
		t.Fail()
	}
	defer conn.Close()
	// 2.创建客户端
	c := pb.NewAddClient(conn)
	// 3.发起gRPC调用
	resp, err := c.Sum(context.Background(), &pb.SumRequest{A: 10, B: 2})
	// 查看结果是否符合预期
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.V, int64(12))
}

// 给Concat方法编写测试
