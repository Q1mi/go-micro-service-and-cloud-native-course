package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	ShortUrlDB ShortUrlDB

	Sequence struct {
		DSN string
	}
}

type ShortUrlDB struct {
	DSN string
}
