package logic

import (
	"context"
	"errors"

	"mall/service/order/api/internal/interceptor"
	"mall/service/order/api/internal/svc"
	"mall/service/order/api/internal/types"
	"mall/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchRequest) (resp *types.SearchResponse, err error) {
	// todo: add your logic here and delete this line
	// 1. 根据请求参数中的订单号查询数据库找到订单记录（课后作业）
	// 2. 根据订单记录中的 user_id 去查询用户数据（通过RPC调用user服务）
	// 假设：user_id = 1676817984
	// 如何存入adminID ？
	l.ctx = context.WithValue(l.ctx, interceptor.CtxKeyAdminID, "33")

	userResp, err := l.svcCtx.UserRPC.GetUser(l.ctx, &userclient.GetUserReq{UserID: 1676817984})
	if err != nil {
		logx.Errorw("UserRPC.GetUser failed", logx.Field("err", err))
		return nil, errors.New("内部错误")
	}
	// 3. 拼接返回结果（因为我们这个接口的数据不是由我一个服务组成的）
	return &types.SearchResponse{
		OrderID:  "1676817984",           // 根据实际查询的订单记录来赋值，这里写的是假数据
		Status:   100,                    // 根据实际查询的订单记录来赋值，这里写的是假数据
		Username: userResp.GetUsername(), // RPC调用user服务拿到的数据
	}, nil
}
