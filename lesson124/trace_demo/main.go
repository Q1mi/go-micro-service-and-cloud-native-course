package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

// jaeger trace Go demo by Q1mi

const (
	serviceName    = "Go-Jaeger-Demo"
	jaegerEndpoint = "127.0.0.1:4318"
)

// setupTracer 设置Tracer
func setupTracer(ctx context.Context) (func(context.Context) error, error) {
	tracerProvider, err := newJaegerTraceProvider(ctx)
	if err != nil {
		return nil, err
	}
	otel.SetTracerProvider(tracerProvider)
	return tracerProvider.Shutdown, nil
}

// newJaegerTraceProvider 创建一个 Jaeger Trace Provider
func newJaegerTraceProvider(ctx context.Context) (*traceSDK.TracerProvider, error) {
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
	traceProvider := traceSDK.NewTracerProvider(
		traceSDK.WithResource(res),
		traceSDK.WithSampler(traceSDK.AlwaysSample()), // 采样
		traceSDK.WithBatcher(exp, traceSDK.WithBatchTimeout(time.Second)),
	)
	return traceProvider, nil
}

// testTracer 编写一个批量创建span的tracer
// trace 和 span
func testTracer(ctx context.Context) {
	tracer := otel.Tracer("test-tracer")
	baseAttrs := []attribute.KeyValue{
		attribute.String("domain", "liwenzhou.com"),
		attribute.Bool("plagiarize", false),
		attribute.Int("code", 7),
	}

	// 开启span
	ctx, span := tracer.Start(ctx, "parent-span", trace.WithAttributes(baseAttrs...))
	// 结束span
	defer span.End()
	// 使用for循环创建多个子span，方便查看效果
	for i := range 10 { // Go1.22+
		// 传入父ctx，开启子span
		_, iSpan := tracer.Start(ctx, fmt.Sprintf("span-%d", i))
		// 随机sleep，模拟子span中耗时的操作
		time.Sleep(time.Duration(rand.Int63n(100)) * time.Millisecond)
		// 子span结束
		iSpan.End()
	}
	fmt.Println("done!")
}

func main() {
	ctx := context.Background()
	shutdown, err := setupTracer(ctx)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = shutdown(ctx)
	}()

	// 批量创建span并上报至Jaeger
	testTracer(ctx)
}

