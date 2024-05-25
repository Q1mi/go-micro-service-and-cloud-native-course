package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	"go.opentelemetry.io/otel/trace"
)

// http client 链路追踪示例
// 上报 trace 数据至 Jaeger

const (
	serviceName     = "httpclient-Demo"
	peerServiceName = "blog"
	jaegerEndpoint  = "127.0.0.1:4318"
	blogURL         = "https://liwenzhou.com"
)

// newJaegerTraceProvider 创建一个 Jaeger Trace Provider
func newJaegerTraceProvider(ctx context.Context) (*sdktrace.TracerProvider, error) {
	// 创建一个使用 HTTP 协议连接本机Jaeger的 Exporter
	exp, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(jaegerEndpoint),
		otlptracehttp.WithInsecure())
	if err != nil {
		return nil, err
	}

	res, err := resource.New(ctx, resource.WithAttributes(semconv.ServiceName(serviceName)))
	if err != nil {
		return nil, err
	}
	// 创建 Provider
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

	otel.SetTracerProvider(tp) // 全局otel设置了tp
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}),
	)
	return tp, nil
}

func main() {
	ctx := context.Background()

	tp, err := initTracer(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	tr := otel.Tracer("http-client") // 直接新建tracer

	// 服务A  --HTTP GET请求-->  服务B(我的博客)
	// 开启一个新的span
	ctx, span := tr.Start(ctx, "GET BLOG", trace.WithAttributes(semconv.PeerService(peerServiceName)))

	// resp, err := http.Get(blogURL) // 不支持传ctx

	// 构造client
	// client := http.Client{Transport: http.DefaultTransport}
	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}

	// 构造请求
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, blogURL, nil)

	// 发请求
	resp, err := client.Do(req)
	span.End()
	if err != nil {
		log.Fatal(err)
	}

	// 解析响应结果
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
	fmt.Printf("body:%s\n", body)

}
