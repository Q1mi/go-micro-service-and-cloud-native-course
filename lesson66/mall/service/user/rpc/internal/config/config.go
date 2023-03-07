package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf

	// mysql
	Mysql struct { // 数据库配置，除mysql外，可能还有mongo等其他数据库
		DataSource string // mysql链接地址，满足 $user:$password@tcp($ip:$port)/$db?$queries 格式即可
	}
	// redis
	CacheRedis cache.CacheConf

	// 引入Consul
	Consul consul.Conf
}
