package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

const (
	serviceName    = "redis-Jaeger-Demo"
	jaegerEndpoint = "127.0.0.1:4318"
)

var tracer = otel.Tracer("redis-demo")

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

func main() {
	ctx := context.Background()

	tp, err := initTracer(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// 启用 tracing
	if err := redisotel.InstrumentTracing(rdb); err != nil {
		panic(err)
	}

	// 启用 metrics
	if err := redisotel.InstrumentMetrics(rdb); err != nil {
		panic(err)
	}

	ctx, span := tracer.Start(ctx, "doSomething")
	defer span.End()

	if err := doSomething(ctx, rdb); err != nil {
		span.RecordError(err) // 记录error
		span.SetStatus(codes.Error, err.Error())
	}
}

// doSomething 使用rdb执行redis操作，mock业务逻辑
func doSomething(ctx context.Context, rdb *redis.Client) error {
	// 写操作
	if err := rdb.Set(ctx, "name", "Q1mi", time.Minute).Err(); err != nil {
		return err
	}
	if err := rdb.Set(ctx, "tag", "OTel", time.Minute).Err(); err != nil {
		return err
	}
	// 读操作
	var wg sync.WaitGroup
	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			val := rdb.Get(ctx, "tag").Val()
			if val != "OTel" {
				log.Printf("%q != %q", val, "OTel")
			}
		}()
	}
	wg.Wait()

	// 删操作
	if err := rdb.Del(ctx, "name").Err(); err != nil {
		return err
	}
	if err := rdb.Del(ctx, "tag").Err(); err != nil {
		return err
	}
	log.Println("done!")
	return nil
}
