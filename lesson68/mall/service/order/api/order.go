package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"mall/service/order/api/internal/config"
	"mall/service/order/api/internal/errorx"
	"mall/service/order/api/internal/handler"
	"mall/service/order/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"

	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

var configFile = flag.String("f", "etc/order-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 注册自定义错误的处理方法
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, any) {
		switch e := err.(type) {
		case errorx.CodeError: // 自定义错误
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
