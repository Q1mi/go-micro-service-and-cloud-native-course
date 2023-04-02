package logic

import (
	"context"

	"shorturl/internal/svc"
	"shorturl/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShorturlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShorturlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShorturlLogic {
	return &ShorturlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShorturlLogic) Shorturl(req *types.Request) (resp *types.Response, err error) {
	// 根据短链接请求的标识符，查找到对应的原始长链接
	// req.ShortURL // 标识符 1ly7vk
	if req.ShortURL == "1ly7vk" {
		return &types.Response{LongURL: "https://www.liwenzhou.com/posts/Go/golang-menu"}, nil
	}
	// 其他查询不到长链接的请求都跳转到 baidu.com
	return &types.Response{LongURL: "https://www.baidu.com"}, nil
}
