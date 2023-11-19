package data

import (
	"context"

	v1 "review-b/api/review/v1"
	"review-b/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type businessRepo struct {
	data *Data
	log  *log.Helper
}

// NewBusinessRepo .
func NewBusinessRepo(data *Data, logger log.Logger) biz.BusinessRepo {
	return &businessRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *businessRepo) Reply(ctx context.Context, param *biz.ReplyParam) (int64, error) {
	r.log.WithContext(ctx).Infof("[data] Reply, param:%v", param)
	// 之前我们都是写操作数据库
	// 而现在我们需要的是通过RPC调用其他的服务
	ret, err := r.data.rc.ReplyReview(ctx, &v1.ReplyReviewRequest{
		ReviewID:  param.ReviewID,
		StoreID:   param.StoreID,
		Content:   param.Content,
		PicInfo:   param.PicInfo,
		VideoInfo: param.VideoInfo,
	})
	r.log.WithContext(ctx).Debugf("ReplyReview return, ret:%v err:%v", ret, err)
	if err != nil {
		return 0, err
	}
	return ret.GetReplyID(), nil
}

func (r *businessRepo) Appeal(ctx context.Context, param *biz.AppealParam) (int64, error) {
	r.log.WithContext(ctx).Infof("[data] Appeal, param:%v", param)
	ret, err := r.data.rc.AppealReview(ctx, &v1.AppealReviewRequest{
		ReviewID:  param.ReviewID,
		StoreID:   param.StoreID,
		Reason:    param.Reason,
		Content:   param.Content,
		PicInfo:   param.PicInfo,
		VideoInfo: param.VideoInfo,
	})
	r.log.WithContext(ctx).Debugf("AppealReview return, ret:%v err:%v", ret, err)
	if err != nil {
		return 0, err
	}
	return ret.GetAppealID(), nil
}
