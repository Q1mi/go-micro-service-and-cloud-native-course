package main

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

const (
	serviceName    = "GORM-Jaeger-Demo"
	jaegerEndpoint = "127.0.0.1:4318"
)

var tracer = otel.Tracer("gorm-demo")

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

	dsn := "root:root1234@tcp(127.0.0.1:13306)/db3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 设置tracing插件
	if err := db.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}

	ctx, span := tracer.Start(ctx, "doSomething")

	if err := doSomething(ctx, db); err != nil {
		span.RecordError(err) // 记录error
		span.SetStatus(codes.Error, err.Error())
	}
	span.End()
}

// doSomething 数据库增删改查操作
func doSomething(ctx context.Context, db *gorm.DB) error {
	type Book struct {
		gorm.Model
		Title string
	}
	// 迁移 schema
	db.AutoMigrate(&Book{})

	// Create
	db.WithContext(ctx).Create(&Book{Title: "《Go语言之路》"})
	// Read
	var book Book
	if err := db.WithContext(ctx).Take(&book).Error; err != nil {
		return err
	}
	// delete
	if err := db.WithContext(ctx).Delete(&book).Error; err != nil {
		return err
	}
	log.Println("done!")
	return nil
}
