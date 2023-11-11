package service

import (
	"context"
	"fmt"

	pb "review-service/api/review/v1"
	"review-service/internal/biz"
	"review-service/internal/data/model"
)

type ReviewService struct {
	pb.UnimplementedReviewServer

	uc *biz.ReviewUsecase
}

func NewReviewService(uc *biz.ReviewUsecase) *ReviewService {
	return &ReviewService{uc: uc}
}

// CreateReview 创建评价
func (s *ReviewService) CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.CreateReviewReply, error) {
	fmt.Printf("[service] CreateReview, req:%#v\n", req)
	// 参数转换
	// 调用biz层
	var anonymous int32
	if req.Anonymous {
		anonymous = 1
	}
	review, err := s.uc.CreateReview(ctx, &model.ReviewInfo{
		UserID:       req.UserID,
		OrderID:      req.OrderID,
		Score:        req.Score,
		ServiceScore: req.Score,
		ExpressScore: req.ExpressScore,
		Content:      req.Content,
		PicInfo:      req.PicInfo,
		VideoInfo:    req.VideoInfo,
		Anonymous:    anonymous,
		Status:       0,
	})
	// 拼装返回结果
	if err != nil {
		return nil, err
	}
	return &pb.CreateReviewReply{ReviewID: review.ReviewID}, nil
}

func (s *ReviewService) GetReview(ctx context.Context, req *pb.GetReviewRequest) (*pb.GetReviewReply, error) {
	fmt.Printf("GetReview req:%#v\n", req)
	review, err := s.uc.GetReview(ctx, req.ReviewID)
	return &pb.GetReviewReply{
		Data: &pb.ReviewInfo{
			ReviewID:     review.ReviewID,
			UserID:       review.UserID,
			OrderID:      review.OrderID,
			Score:        review.Score,
			ServiceScore: review.ServiceScore,
			ExpressScore: review.ExpressScore,
			Content:      review.Content,
			PicInfo:      review.PicInfo,
			VideoInfo:    review.VideoInfo,
			Status:       review.Status,
		},
	}, err
}
func (s *ReviewService) AuditReview(ctx context.Context, req *pb.AuditReviewRequest) (*pb.AuditReviewReply, error) {
	fmt.Printf("AuditReview req:%#v\n", req)
	err := s.uc.AuditReview(ctx, &biz.AuditParam{
		ReviewID:  req.GetReviewID(),
		OpUser:    req.GetOpUser(),
		OpReason:  req.GetOpReason(),
		OpRemarks: req.GetOpRemarks(),
		Status:    req.GetStatus(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.AuditReviewReply{
		ReviewID: req.ReviewID,
		Status:   req.Status,
	}, nil
}

// ReplyReview 回复评价
func (s *ReviewService) ReplyReview(ctx context.Context, req *pb.ReplyReviewRequest) (*pb.ReplyReviewReply, error) {
	fmt.Printf("[service] ReplyReview req:%#v\n", req)
	// 调用biz层
	reply, err := s.uc.CreateReply(ctx, &biz.ReplyParam{
		ReviewID:  req.GetReviewID(),
		StoreID:   req.GetStoreID(),
		Content:   req.GetContent(),
		PicInfo:   req.GetPicInfo(),
		VideoInfo: req.GetVideoInfo(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.ReplyReviewReply{ReplyID: reply.ReplyID}, nil
}
func (s *ReviewService) AppealReview(ctx context.Context, req *pb.AppealReviewRequest) (*pb.AppealReviewReply, error) {
	return &pb.AppealReviewReply{}, nil
}
func (s *ReviewService) AuditAppeal(ctx context.Context, req *pb.AuditAppealRequest) (*pb.AuditAppealReply, error) {
	return &pb.AuditAppealReply{}, nil
}
func (s *ReviewService) ListReviewByUserID(ctx context.Context, req *pb.ListReviewByUserIDRequest) (*pb.ListReviewByUserIDReply, error) {
	return &pb.ListReviewByUserIDReply{}, nil
}
