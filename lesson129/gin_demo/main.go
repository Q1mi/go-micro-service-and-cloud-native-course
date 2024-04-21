package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// gin 框架链路追踪示例
// 上报 trace 数据至Jaeger

const (
	serviceName    = "Gin-Jaeger-Demo"
	jaegerEndpoint = "127.0.0.1:4318"
)

var tracer = otel.Tracer("gin-server")

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
		log.Fatalf("initTracer failed, err:%v", err)
	}
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatalf("tp.Shutdown failed, err:%v", err)
		}
	}()

	// 初始化zap日志对象
	logger := otelzap.New(
		zap.NewExample(),                    // zap实例，按需配置
		otelzap.WithMinLevel(zap.InfoLevel), // 指定日志级别
		otelzap.WithTraceIDField(true),      // 在日志中记录 traceID
	)
	defer logger.Sync()

	// 替换全局
	otelzap.ReplaceGlobals(logger)

	r := gin.New()

	// 注册trace中间件
	// 设置 otelgin 中间件
	r.Use(otelgin.Middleware(serviceName))

	// 在响应头里记录traceId
	r.Use(func(c *gin.Context) {
		// 从原始请求的ctx中拿到span对象 trace.SpanFromContext(c.Request.Context())
		span := trace.SpanFromContext(c.Request.Context())
		// 从span对象取出traceID对象
		traceID := span.SpanContext().TraceID().String()
		c.Header("Trace-Id", traceID) // 写响应头
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		name := getUser(c, id)
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"id":   id,
		})
	})
	_ = r.Run(":8080")
}

func getUser(c *gin.Context, id string) string {
	ctx, span := tracer.Start(
		c.Request.Context(),
		"getUser",
		trace.WithAttributes(attribute.String("id", id)), // 添加自定义的属性
	)
	defer span.End()

	// 在业务逻辑代码中记录日志（包含traceID）
	otelzap.Ctx(ctx).Info("getUser logic", zap.String("id", id))

	// mock 业务逻辑
	if id == "7" {
		return "Q1mi"
	}
	return "unknown"
}
