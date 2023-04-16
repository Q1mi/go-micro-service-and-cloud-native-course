package svc

import (
	"shortener/internal/config"
	"shortener/model"
	"shortener/sequence"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	ShortUrlModel model.ShortUrlMapModel // short_url_map

	Sequence sequence.Sequence // sequence
	// Sequence *sequence.Redis // sequence
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.ShortUrlDB.DSN)
	return &ServiceContext{
		Config:        c,
		ShortUrlModel: model.NewShortUrlMapModel(conn),
		Sequence:      sequence.NewMySQL(c.Sequence.DSN), // sequence
		// Sequence:      sequence.Newedis(redisAddr), // sequence
	}
}
