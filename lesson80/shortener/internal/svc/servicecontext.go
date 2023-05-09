package svc

import (
	"shortener/internal/config"
	"shortener/model"
	"shortener/sequence"

	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	ShortUrlModel model.ShortUrlMapModel // short_url_map

	Sequence sequence.Sequence // sequence
	// Sequence *sequence.Redis // sequence

	ShortUrlBlackList map[string]struct{}

	// bloom filter
	Filter *bloom.Filter
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.ShortUrlDB.DSN)
	// 把配置文件中配置的黑名单加载到map，方便后续判断
	m := make(map[string]struct{}, len(c.ShortUrlBlackList))
	for _, v := range c.ShortUrlBlackList {
		m[v] = struct{}{}
	}
	// 初始化布隆过滤器
	// 初始化 redisBitSet
	store := redis.New(c.CacheRedis[0].Host, func(r *redis.Redis) {
		r.Type = redis.NodeType
	})
	// 声明一个bitSet, key="test_key"名且bits是1024位
	filter := bloom.New(store, "bloom_filter", 20*(1<<20))
	// 加载已有的短链接数据

	return &ServiceContext{
		Config:        c,
		ShortUrlModel: model.NewShortUrlMapModel(conn, c.CacheRedis),
		Sequence:      sequence.NewMySQL(c.Sequence.DSN), // sequence
		// Sequence:      sequence.Newedis(redisAddr), // sequence
		ShortUrlBlackList: m, // 短链接黑名单map
		Filter:            filter,
	}
}

/*
// 注意导入的是这个bloom
import github.com/bits-and-blooms/bloom/v3

// 初始化bloomfilter
filter := bloom.NewWithEstimates(1<<20, 0.01)

// loadDataToBloomFilter 加载已有的短链接数据至布隆过滤器
func loadDataToBloomFilter(conn sqlx.SqlConn, filter *bloom.BloomFilter) error {
	// 循环从数据库查询数据加载至filter
	if conn == nil || filter == nil {
		return errors.New("loadDataToBloomFilter invalid param")
	}

	// 查总数
	total := 0
	if err := conn.QueryRow(&total, "select count(*) from short_url_map where is_del=0"); err != nil {
		logx.Errorw("conn.QueryRowCount failed", logx.LogField{Key: "err", Value: err.Error()})
		return err
	}
	logx.Infow("total data", logx.LogField{Key: "total", Value: total})
	if total == 0 {
		logx.Info("no data need to load")
		return nil
	}
	pageTotal := 0
	pageSize := 20
	if total%pageSize == 0 {
		pageTotal = total / pageSize
	} else {
		pageTotal = total/pageSize + 1
	}
	logx.Infow("pageTotal", logx.LogField{Key: "pageTotal", Value: pageTotal})
	// 循环查询所有数据
	for page := 1; page <= pageTotal; page++ {
		offset := (page - 1) * pageSize
		surls := []string{}
		if err := conn.QueryRows(&surls, "select surl from short_url_map where is_del=0 limit ?,?", offset, pageSize); err != nil {
			return err
		}

		for _, surl := range surls {
			filter.AddString(surl)
		}
	}
	logx.Info("load data to bloom success")
	return nil
}

*/
