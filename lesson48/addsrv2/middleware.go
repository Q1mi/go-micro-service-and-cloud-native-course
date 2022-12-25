package main

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
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
