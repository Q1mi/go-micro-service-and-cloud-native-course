package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"shortener/internal/svc"
	"shortener/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	Err404 = errors.New("404")
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 自己写缓存          surl -> lurl
// go-zero自带的缓存   surl -> 数据行

func (l *ShowLogic) Show(req *types.ShowRequest) (resp *types.ShowResponse, err error) {
	// 查看短链接，输入 q1mi.cn/lusytc -> 重定向到真实的链接
	// req.ShortUrl = lusytc
	// 1. 根据短链接查询原始的长链接
	// 1.0 布隆过滤器
	// 不存在的短链接直接返回404，不需要后续处理
	// a.基于内存版本，服务重启之后就没了，所以每次重启都要重新加载一下已有的短链接（从数据库查询）
	// b.基于Redis版本，go-zero自带：https://go-zero.dev/cn/docs/blog/governance/bloom/
	exist, err := l.svcCtx.Filter.Exists([]byte(req.ShortUrl))
	if err != nil {
		logx.Errorw("Bloom Filter failed", logx.LogField{Value: err.Error(), Key: "err"})
	}
	// 不存在的短链接直接返回
	if !exist {
		return nil, Err404
	}
	fmt.Println("开始查询缓存和DB...")
	// 1.1 查询数据库之前可增加缓存层
	// go-zero的缓存支持singleflight
	u, err := l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{Valid: true, String: req.ShortUrl})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, Err404
		}
		logx.Errorw("ShortUrlModel.FindOneBySurl failed", logx.LogField{Value: err.Error(), Key: "err"})
		return nil, err
	}
	// 2. 返回查询到的长链接，在调用handler层返回重定向响应
	return &types.ShowResponse{LongUrl: u.Lurl.String}, nil
}
