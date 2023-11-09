package service

import (
	"context"
	"review-b/internal/biz"

	pb "review-b/api/business/v1"
)

type BusinessService struct {
	pb.UnimplementedBusinessServer

	uc *biz.BusinessUsecase
}

func NewBusinessService(uc *biz.BusinessUsecase) *BusinessService {
	return &BusinessService{uc: uc}
}

func (s *BusinessService) ReplyReview(ctx context.Context, req *pb.ReplyReviewRequest) (*pb.ReplyReviewReply, error) {
	return &pb.ReplyReviewReply{}, nil
}
