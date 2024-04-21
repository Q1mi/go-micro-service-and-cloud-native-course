package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"pb"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"google.golang.org/grpc"
)

const (
	serviceName       = "gRPC-Jaeger-Demo"
	jaegerRPCEndpoint = "127.0.0.1:4317" // Jaeger RPC端口
)

var tracer = otel.Tracer("grpc-server-example")

// newJaegerTraceProvider 创建一个 Jaeger Trace Provider
func newJaegerTraceProvider(ctx context.Context) (*sdktrace.TracerProvider, error) {
	// 创建一个使用 HTTP 协议连接本机Jaeger的 Exporter
	exp, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint(jaegerRPCEndpoint),
		otlptracegrpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	res, err := resource.New(ctx, resource.WithAttributes(semconv.ServiceName(serviceName)))
	if err != nil {
		return nil, err
	}
	traceProvider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.AlwaysSample()), // 采样
		sdktrace.WithBatcher(exp, sdktrace.WithBatchTimeout(time.Second)),
	)
	return traceProvider, nil
}

// initTracer 初始化 Tracer
func initTracer(ctx context.Context) (*sdktrace.TracerProvider, error) {
	tp, err := newJaegerTraceProvider(ctx)
	if err != nil {
		return nil, err
	}

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}),
	)
	return tp, nil
}

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	time.Sleep(time.Duration(rand.Int31n(100)) * time.Millisecond) // 随机sleep
	return &pb.HelloResponse{Reply: "Hello " + in.Name}, nil
}

func main() {
	ctx := context.Background()
	tp, err := initTracer(ctx)
	if err != nil {
		log.Fatalf("initTracer failed, err:%v", err)
	}
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatalf("tp.Shutdown failed, err:%v", err)
		}
	}()

	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	// 创建gRPC服务器
	s := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()), // 设置 StatsHandler
	)

	pb.RegisterGreeterServer(s, &server{}) // 在gRPC服务端注册服务
	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
