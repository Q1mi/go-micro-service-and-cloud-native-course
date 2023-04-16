package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"shortener/internal/svc"
	"shortener/internal/types"
	"shortener/model"
	"shortener/pkg/base62"
	"shortener/pkg/connect"
	"shortener/pkg/md5"
	"shortener/pkg/urltool"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Convert 转链业务逻辑：输入一个长链接 --> 转为短链接
func (l *ConvertLogic) Convert(req *types.ConvertRequest) (resp *types.ConvertResponse, err error) {
	// 1. 校验输入的数据
	// 1.1 数据不能空
	// if len(req.LongUrl) == 0{}
	// 使用validator包来做参数校验
	// 1.2 输入的长链接必须是一个能请求通的网址
	if ok := connect.Get(req.LongUrl); !ok { // xxx/123as
		return nil, errors.New("无效的链接")
	}
	// 1.3 判断之前是否已经转链过（数据库中是否已存在该长链接）
	// 1.3.1 给长链接生成md5
	md5Value := md5.Sum([]byte(req.LongUrl)) // 注意！注意！注意！ 这里使用的是项目中封装的 pkg/md5 包
	// 1.3.2 拿md5去数据库中查是否存在
	u, err := l.svcCtx.ShortUrlModel.FindOneByMd5(l.ctx, sql.NullString{String: md5Value, Valid: true})
	if err != sqlx.ErrNotFound {
		if err == nil {
			return nil, fmt.Errorf("该链接已被转为%s", u.Surl.String)
		}
		logx.Errorw("ShortUrlModel.FindOneByMd5 failed", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}
	// 1.4 输入的不能是一个短链接（避免循环转链）
	// 输入的是一个完整的url q1mi.cn/1d12a?name=q1mi
	basePath, err := urltool.GetBasePath(req.LongUrl)
	if err != nil {
		logx.Errorw("urltool.GetBasePath failed", logx.LogField{Key: "lurl", Value: req.LongUrl}, logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}
	_, err = l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{String: basePath, Valid: true})
	if err != sqlx.ErrNotFound {
		if err == nil {
			return nil, errors.New("该链接已经是短链了")
		}
		logx.Errorw("ShortUrlModel.FindOneBySurl failed", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}

	var short string
	for {
		// 2. 取号 基于MySQL实现的发号器
		// 每来一个转链请求，我们就使用 REPLACE INTO语句往 sequence 表插入一条数据，并且取出主键id作为号码
		seq, err := l.svcCtx.Sequence.Next()
		if err != nil {
			logx.Errorw("Sequence.Next() failed", logx.LogField{Key: "err", Value: err.Error()})
			return nil, err
		}
		fmt.Println(seq) // 1
		// 3. 号码转短链
		// 3.1 安全性  1En = 6347
		short = base62.Int2String(seq)
		// 3.2 短域名黑名单避免某些特殊词比如 api、health、fuck等等
		if _, ok := l.svcCtx.ShortUrlBlackList[short]; !ok {
			break // 生成不在黑名单里的短链接就跳出for循环
		}
	}

	fmt.Printf("short:%v\n", short)
	// 4. 存储长短链接映射关系
	if _, err := l.svcCtx.ShortUrlModel.Insert(
		l.ctx,
		&model.ShortUrlMap{
			Lurl: sql.NullString{String: req.LongUrl, Valid: true},
			Md5:  sql.NullString{String: md5Value, Valid: true},
			Surl: sql.NullString{String: short, Valid: true},
		},
	); err != nil {
		logx.Errorw("ShortUrlModel.Insert() failed", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}
	// 5. 返回响应
	// 5.1 返回的是 短域名+短链接  q1mi.cn/1En
	shortUrl := l.svcCtx.Config.ShortDoamin + "/" + short
	return &types.ConvertResponse{ShortUrl: shortUrl}, nil
}
