package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// BusinessRepo is a Greater repo.
type BusinessRepo interface {
	Save(context.Context) error
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
