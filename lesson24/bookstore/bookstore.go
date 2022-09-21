package main

import (
	"bookstore/pb"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

// bookstore grpc服务

type server struct {
	pb.UnimplementedBookstoreServer

	bs *bookstore // data.go
}

// ListShelves 列出所有书架的RPC方法
func (s *server) ListShelves(ctx context.Context, in *emptypb.Empty) (*pb.ListShelvesResponse, error) {
	// 调用orm操作的那些方法
	sl, err := s.bs.ListShelves(ctx)
	if err == gorm.ErrEmptySlice { // 没有数据
		return &pb.ListShelvesResponse{}, nil
	}
	if err != nil { // 查询数据库失败
		return nil, status.Error(codes.Internal, "query failed")
	}
	// 封装返回数据
	nsl := make([]*pb.Shelf, 0, len(sl))
	for _, s := range sl {
		nsl = append(nsl, &pb.Shelf{
			Id:    s.ID,
			Theme: s.Theme,
			Size:  s.Size,
		})
	}
	return &pb.ListShelvesResponse{Shelves: nsl}, nil
}

// CreateShelf 创建书架
func (s *server) CreateShelf(ctx context.Context, in *pb.CreateShelfRequest) (*pb.Shelf, error) {
	// 参数检查
	if len(in.GetShelf().GetTheme()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid theme")
	}
	// 准备数据
	data := Shelf{
		Theme: in.GetShelf().GetTheme(),
		Size:  in.GetShelf().GetSize(),
	}
	// 去数据库创建
	ns, err := s.bs.CreateShelf(ctx, data)
	if err != nil {
		return nil, status.Error(codes.Internal, "create failed")
	}
	return &pb.Shelf{Id: ns.ID, Theme: ns.Theme, Size: ns.Size}, nil
}

// GetShelf 根据id获取书架
func (s *server) GetShelf(ctx context.Context, in *pb.GetShelfRequest) (*pb.Shelf, error) {
	// 参数check
	if in.GetShelf() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid shelf id")
	}
	// 查询数据库
	shelf, err := s.bs.GetShelf(ctx, in.GetShelf())
	if err != nil {
		return nil, status.Error(codes.Internal, "query failed")
	}
	// 封装返回数据
	return &pb.Shelf{Id: shelf.ID, Theme: shelf.Theme, Size: shelf.Size}, nil
}

// DeleteShelf 根据ID删除书架
func (s *server) DeleteShelf(ctx context.Context, in *pb.DeleteShelfRequest) (*emptypb.Empty, error) {
	// 参数check
	if in.GetShelf() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid shelf id")
	}
	err := s.bs.DeleteShelf(ctx, in.GetShelf())
	if err != nil {
		return nil, status.Error(codes.Internal, "delete failed")
	}
	return &emptypb.Empty{}, nil
}
