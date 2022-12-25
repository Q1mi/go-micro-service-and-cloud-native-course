package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/log"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"addsrv/pb"
)

// go-kit addService demo 3

var (
	httpAddr = flag.Int("http-addr", 8080, "HTTP端口")
	gRPCAddr = flag.Int("grpc-addr", 8972, "gRPC端口")
)

func main() {
	// 前置资源初始化

	srv := NewService()
	// 初始化带logger的service
	logger := log.NewJSONLogger(os.Stdout)
	srv = NewLogMiddleware(logger, srv)
	// metrics
	// instrumentation
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "add_srv",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "add_srv",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "add_srv",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	// metrics 中间件
	srv = instrumentingMiddleware{
		requestCount:   requestCount,
		requestLatency: requestLatency,
		countResult:    countResult,
		next:           srv,
	}

	var g errgroup.Group

	// HTTP
	g.Go(func() error {
		httpListener, err := net.Listen("tcp", fmt.Sprintf(":%d", *httpAddr))
		if err != nil {
			fmt.Printf("net.Listen %d faield, err:%v\n", *httpAddr, err)
			return err
		}
		defer httpListener.Close()
		// 初始化go-kit logger
		logger := log.NewLogfmtLogger(os.Stderr)
		httpHandler := NewHTTPServer(srv, logger)

		// http
		// http.Handle("/metrics", promhttp.Handler())

		// gin
		httpHandler.(*gin.Engine).GET("/metrics", gin.WrapH(promhttp.Handler()))

		return http.Serve(httpListener, httpHandler)
	})
	// gRPC
	g.Go(func() error {
		grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", *gRPCAddr))
		if err != nil {
			fmt.Printf("net.Listen %d faield, err:%v\n", *gRPCAddr, err)
			return err
		}
		defer grpcListener.Close()

		s := grpc.NewServer() // gRPC Server
		pb.RegisterAddServer(s, NewGRPCServer(srv))
		return s.Serve(grpcListener)
	})

	// wait
	if err := g.Wait(); err != nil {
		fmt.Printf("server exit with error:%v\n", err)
	}
}
