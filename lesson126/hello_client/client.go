package main

import (
	"context"
	"flag"
	"log"
	"time"

	"pb"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const (
	serviceName    = "gRPC-Jaeger-Demo"
	jaegerEndpoint = "127.0.0.1:4317"
)

var tracer = otel.Tracer("grpc-client-example")

// newJaegerTraceProvider 创建一个 Jaeger Trace Provider
func newJaegerTraceProvider(ctx context.Context) (*sdktrace.TracerProvider, error) {
	// 创建一个使用 HTTP 协议连接本机Jaeger的 Exporter
	exp, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint(jaegerEndpoint),
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

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "127.0.0.1:8972", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()

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

	// 连接到server端
	conn, err := grpc.NewClient(
		*addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() { _ = conn.Close() }()

	c := pb.NewGreeterClient(conn)

	// 执行RPC调用并打印收到的响应数据
	md := metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
		"client-id", "hello-client-q1mi",
		"user-id", "7",
	)
	ctx = metadata.NewOutgoingContext(context.Background(), md)

	ctx, span := tracer.Start(ctx, "c.SayHello", trace.WithAttributes(attribute.String("name", "gRPC-client")))
	defer span.End()
	// RPC调用
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetReply())
}
