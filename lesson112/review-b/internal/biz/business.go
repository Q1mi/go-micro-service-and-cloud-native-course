package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type ReplyParam struct {
	ReviewID  int64
	StoreID   int64
	Content   string
	PicInfo   string
	VideoInfo string
}

type AppealParam struct {
	ReviewID  int64
	StoreID   int64
	Reason    string
	Content   string
	PicInfo   string
	VideoInfo string
}

// BusinessRepo is a Greater repo.
type BusinessRepo interface {
	Reply(context.Context, *ReplyParam) (int64, error)
	Appeal(context.Context, *AppealParam) (int64, error)
}

// BusinessUsecase is a Greeter usecase.
type BusinessUsecase struct {
	repo BusinessRepo
	log  *log.Helper
}

// NewBusinessUsecase new a Greeter usecase.
func NewBusinessUsecase(repo BusinessRepo, logger log.Logger) *BusinessUsecase {
	return &BusinessUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateReply 创建回复
// service层调用此方法
func (uc *BusinessUsecase) CreateReply(ctx context.Context, param *ReplyParam) (int64, error) {
	uc.log.WithContext(ctx).Infof("[biz] CreateReply param:%v", param)
	return uc.repo.Reply(ctx, param)
}

// CreateAppeal 创建申诉
func (uc *BusinessUsecase) CreateAppeal(ctx context.Context, param *AppealParam) (int64, error) {
	uc.log.WithContext(ctx).Infof("[biz] CreateReply param:%v", param)
	return uc.repo.Appeal(ctx, param)
}
