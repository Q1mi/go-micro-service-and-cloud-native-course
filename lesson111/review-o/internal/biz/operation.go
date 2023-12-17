package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// AuditReviewParam 审核评价的参数
type AuditReviewParam struct {
	ReviewID  int64
	Status    int
	OpReason  string
	OpRemarks string
	OpUser    string
}

// AuditAppealParam 审核申诉的参数
type AuditAppealParam struct {
	AppealID  int64
	ReviewID  int64
	StoreID   int64
	Status    int
	OpReason  string
	OpRemarks string
	OpUser    string
}

// OperationRepo is a Operation repo.
type OperationRepo interface {
	AuditReview(context.Context, *AuditReviewParam) error
	AuditAppeal(context.Context, *AuditAppealParam) error
}

// OperationUsecase is a Operation usecase.
type OperationUsecase struct {
	repo OperationRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewOperationUsecase(repo OperationRepo, logger log.Logger) *OperationUsecase {
	return &OperationUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *OperationUsecase) AuditReview(ctx context.Context, param *AuditReviewParam) error {
	uc.log.WithContext(ctx).Infof("AuditReview，param:%v", param)
	return uc.repo.AuditReview(ctx, param)
}

func (uc *OperationUsecase) AuditAppeal(ctx context.Context, param *AuditAppealParam) error {
	uc.log.WithContext(ctx).Infof("AuditAppeal,param:%v", param)
	return uc.repo.AuditAppeal(ctx, param)
}
