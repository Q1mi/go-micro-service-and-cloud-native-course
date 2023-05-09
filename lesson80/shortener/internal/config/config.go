package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	ShortUrlDB ShortUrlDB

	Sequence struct {
		DSN string
	}

	BaseString string // bas62指定基础字符串

	ShortUrlBlackList []string
	ShortDoamin       string

	CacheRedis cache.CacheConf // redis缓存
}

type ShortUrlDB struct {
	DSN string
}
