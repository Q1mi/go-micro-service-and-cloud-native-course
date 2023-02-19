package logic

import (
	"context"
	"errors"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	// todo: add your logic here and delete this line
	// 遇事不决先写注释！
	// 1. 拿到请求参数
	// req.UserID
	// 2. 根据用户id查询数据库
	user, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, req.UserID)
	if err != nil { // 1.查数据库本身失败了;2.没查询到ErrNotFound
		if err != sqlx.ErrNotFound {
			logx.Errorw("UserModel.FindOneByUserId failed", logx.Field("err", err))
			return nil, errors.New("内部错误")
		}
		return nil, errors.New("用户不存在")
	}
	// 3. 格式化数据(数据库里存的字段和前端要求的字段不太一致)
	// 4. 返回响应
	return &types.DetailResponse{
		Username: user.Username,
		Gender:   int(user.Gender),
	}, nil
}
