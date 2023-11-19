package data

import (
	"context"

	reviewv1 "review-o/api/review/v1"
	"review-o/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type operationRepo struct {
	data *Data
	log  *log.Helper
}

// NewOperationRepo .
func NewOperationRepo(data *Data, logger log.Logger) biz.OperationRepo {
	return &operationRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *operationRepo) AuditReview(ctx context.Context, param *biz.AuditReviewParam) error {
	r.log.WithContext(ctx).Infof("AuditReview, param:%v", param)
	ret, err := r.data.rc.AuditReview(ctx, &reviewv1.AuditReviewRequest{
		ReviewID:  param.ReviewID,
		Status:    int32(param.Status),
		OpUser:    param.OpUser,
		OpReason:  param.OpReason,
		OpRemarks: &param.OpRemarks,
	})
	r.log.WithContext(ctx).Debugf("AuditReview reply ret: %v, err:%v", ret, err)
	return err
}

func (r *operationRepo) AuditAppeal(ctx context.Context, param *biz.AuditAppealParam) error {
	r.log.WithContext(ctx).Infof("AuditReview, param:%v", param)
	ret, err := r.data.rc.AuditAppeal(ctx, &reviewv1.AuditAppealRequest{
		AppealID:  param.AppealID,
		ReviewID:  param.ReviewID,
		Status:    int32(param.Status),
		OpUser:    param.OpUser,
		OpRemarks: &param.OpRemarks,
	})
	r.log.WithContext(ctx).Debugf("AuditReview reply ret: %v, err:%v", ret, err)
	return err
}
