package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	ShortUrlDB ShortUrlDB

	Sequence struct {
		DSN string
	}

	BaseString string // bas62指定基础字符串

	ShortUrlBlackList []string
	ShortDoamin       string
}

type ShortUrlDB struct {
	DSN string
}
