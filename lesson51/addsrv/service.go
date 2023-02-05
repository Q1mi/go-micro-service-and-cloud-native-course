package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/log"
	"google.golang.org/grpc"

	sdconsul "github.com/go-kit/kit/sd/consul"
	consulapi "github.com/hashicorp/consul/api"
)

// service层
// 所有跟业务逻辑相关的我们都应该放在这一层。

// 1.1 业务逻辑抽象为接口

type AddService interface {
	Sum(ctx context.Context, a, b int) (int, error)
	Concat(ctx context.Context, a, b string) (string, error)
}

// 1.2 实现接口

// addService 一个AddService接口的具体实现
// 它的内部可以按需添加各种字段
type addService struct {
	// db db.Conn
	// logger zap.Logger
}

var (
	// ErrEmptyString 两个参数都是空字符串
	ErrEmptyString = errors.New("两个参数都是空字符串")
)

// Sum 返回两个数的和
func (s addService) Sum(_ context.Context, a, b int) (int, error) {
	// 业务逻辑
	// 1.查询数据
	// s.db.Query()

	// 2.处理数据
	return a + b, nil
}

// Concat 拼接两个字符串
func (s addService) Concat(_ context.Context, a, b string) (string, error) {
	if a == "" && b == "" {
		return "", ErrEmptyString
	}
	// 记录日志 ？？？
	// zap.L().Error()  // 引用全局变量方式
	// s.logger.Error()  // 结构体嵌入logger
	return a + b, nil
}

// NewService addService的构造函数
func NewService() AddService {
	return &addService{
		// db:db
	}
}

// logMiddleware

type logMiddleware struct {
	logger log.Logger
	next   AddService //  嵌入接口
}

func NewLogMiddleware(logger log.Logger, svc AddService) AddService {
	return &logMiddleware{
		logger: logger,
		next:   svc,
	}
}

func (s logMiddleware) Sum(ctx context.Context, a, b int) (res int, err error) {
	defer func(start time.Time) {
		s.logger.Log(
			"method", "sum",
			"a", a,
			"b", b,
			"res", res,
			"err", err,
			"cast", time.Since(start),
		) // 记录日志
	}(time.Now())

	res, err = s.next.Sum(ctx, a, b) // 业务逻辑
	return
}

func (s logMiddleware) Concat(ctx context.Context, a, b string) (res string, err error) {
	defer func(start time.Time) {
		s.logger.Log(
			"method", "concat",
			"a", a,
			"b", b,
			"res", res,
			"err", err,
			"cast", time.Since(start),
		) // 记录日志
	}(time.Now())
	res, err = s.next.Concat(ctx, a, b)
	return
}

// metrics
type instrumentingMiddleware struct {
	requestCount   metrics.Counter // 计数器
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           AddService
}

// func NewInstrumentingMiddleware(counter metrics.Counter, ) AddService {
// 	return &instrumentingMiddleware{}
// }

func (im instrumentingMiddleware) Sum(ctx context.Context, a, b int) (res int, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "sum", "error", fmt.Sprint(err != nil)}
		im.requestCount.With(lvs...).Add(1)
		im.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		im.countResult.Observe(float64(res))
	}(time.Now())
	res, err = im.next.Sum(ctx, a, b)
	return
}

func (im instrumentingMiddleware) Concat(ctx context.Context, a, b string) (res string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "concat", "error", "false"}
		im.requestCount.With(lvs...).Add(1)
		im.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	res, err = im.next.Concat(ctx, a, b)
	return
}

// trim相关
type withTrimMiddleware struct {
	next        AddService
	trimService endpoint.Endpoint // trim 交给这个endpoint处理
}

func NewServiceWithTrim(trimEndpoint endpoint.Endpoint, svc AddService) AddService {
	return &withTrimMiddleware{
		trimService: trimEndpoint,
		next:        svc,
	}
}

// 为 withTrimMiddleware 实现 AddService 接口
func (tm withTrimMiddleware) Sum(ctx context.Context, a, b int) (res int, err error) {
	return tm.next.Sum(ctx, a, b) // 复用之前的逻辑
}

func (tm withTrimMiddleware) Concat(ctx context.Context, a, b string) (res string, err error) {
	// 需要新的逻辑处理
	// 外部调用我们的Concat方法时
	// 1. 发起RPC调用 trim_service 对数据进行处理 （调用其他服务/依赖其他的服务）
	respA, err := tm.trimService(ctx, trimRequest{s: a}) // 执行，其实是作为客户端对外发起请求
	if err != nil {
		return "", err
	}
	respB, err := tm.trimService(ctx, trimRequest{s: b}) // 执行，其实是作为客户端对外发起请求
	if err != nil {
		return "", err
	}
	trimA := respA.(trimResponse) // 拿到处理后的响应
	trimB := respB.(trimResponse) // 拿到处理后的响应

	// 2. 拿到处理后的数据再拼接
	return tm.next.Concat(ctx, trimA.s, trimB.s)
}

// consul
// 从注册中心获取trim服务的地址
// 基于consul实现对trim service的服务发现

func getTrimServiceFromConsul(consulAddr string, logger log.Logger, srvName string, tags []string) (endpoint.Endpoint, error) {
	// 1. 连consul
	cfg := consulapi.DefaultConfig()
	cfg.Address = consulAddr
	cc, err := consulapi.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	// 2. 使用go kit 提供的适配器
	sdClient := sdconsul.NewClient(cc)

	instancer := sdconsul.NewInstancer(sdClient, logger, srvName, tags, true)
	// 3. Endpointer
	endpointer := sd.NewEndpointer(instancer, factory, logger)
	// 4. Balancer
	balancer := lb.NewRoundRobin(endpointer)
	// 5. retry
	retry := lb.Retry(3, time.Second, balancer)
	return retry, nil
}

func factory(instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc.Dial(instance, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	e := makeTrimEndpoint(conn)
	return e, conn, err
}
