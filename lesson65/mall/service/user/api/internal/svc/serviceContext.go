package svc

import (
	"api/internal/config"
	"api/internal/middleware"
	"mall/service/user/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	Cost      rest.Middleware // 自定义路由中间件,字段名要与.api文件中的声明一致
	UserModel model.UserModel // 加入User表增删改查操作Model
}

func NewServiceContext(c config.Config) *ServiceContext {
	// UserModel -> 接口类型
	// *defaultUserModel 实现了接口
	// 调用构造函数得到 *model.defaultUserModel
	// NewUserModel(conn sqlx.SqlConn)
	// 需要 sqlx.SqlCon 的数据库连接

	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlxConn, c.CacheRedis),
		Cost:      middleware.NewCostMiddleware().Handle,
	}
}
