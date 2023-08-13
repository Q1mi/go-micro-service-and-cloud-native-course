package server

import (
	v1 "bubble/api/bubble/v1"
	"bubble/internal/conf"
	"bubble/internal/service"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"

	// "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// Middleware 自定义中间件
// type middleware func(Handler) Handler
// type Handler func(ctx context.Context, req interface{}) (interface{}, error)
func Middleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			// 执行之前做点事
			fmt.Println("Middleware：执行handle之前")
			// 做token校验
			if tr, ok := transport.FromServerContext(ctx); ok {
				token := tr.RequestHeader().Get("token")
				fmt.Printf("token:%v\n", token)
			}
			// transport.FromClientContext(ctx)  // 客户端用
			defer func() {
				fmt.Println("Middleware：执行handle之后")
			}()
			return handler(ctx, req) // 执行目标handler
		}
	}
}

// Middleware1 自定义中间件
func Middleware1(opts ...string) middleware.Middleware {
	return func(middleware.Handler) middleware.Handler {
		// opts
		return nil
	}
}

// Middleware2 自定义中间件2，相比Middleware1 失去了一些灵活性
func Middleware2(middleware.Handler) middleware.Handler {
	return nil
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, todo *service.TodoService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			// recovery.Recovery(),
			// Middleware(),
			// // Middleware1("a"),
			// // Middleware1("b"),
			// // Middleware1("c"),
			// // Middleware2,
			// // jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
			// // 	return []byte("123"), nil
			// // }),
			recovery.Recovery(), // 全局中间件
			selector.Server( // 特定Path才执行的中间件
				Middleware(),
			).
				Path("/api.bubble.v1.Todo/CreateTodo").
				Build(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	// 替换默认的HTTP响应编码器
	opts = append(opts, http.ResponseEncoder(responseEncoder))
	// 替换默认的错误响应编码器
	opts = append(opts, http.ErrorEncoder(errorEncoder))

	srv := http.NewServer(opts...)
	v1.RegisterTodoHTTPServer(srv, todo)
	return srv
}
