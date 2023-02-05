package main

import (
	"context"
	"errors"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"golang.org/x/time/rate"
)

// loggingMiddleware 日志中间件
func loggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "开始调用")
			start := time.Now()
			defer logger.Log("msg", "调用结束", "cast", time.Since(start))
			return next(ctx, request)
		}
	}
}

// ErrRateLimit 请求速率限制
var ErrRateLimit = errors.New("request rate limit")

// rateMiddleware 限流中间件
// "golang.org/x/time/rate"
func rateMiddleware(limit *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			// 限流的逻辑
			if limit.Allow() {
				return next(ctx, request)
			} else {
				return nil, ErrRateLimit
			}
		}
	}
}
