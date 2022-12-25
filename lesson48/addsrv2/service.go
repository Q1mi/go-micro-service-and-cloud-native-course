package main

import (
	"context"
	"errors"
	"time"

	"github.com/go-kit/log"
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
