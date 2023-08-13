package service

import (
	"context"
	"errors"

	pb "bubble/api/bubble/v1"
	v1 "bubble/api/bubble/v1"
	"bubble/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type TodoService struct {
	pb.UnimplementedTodoServer

	// 嵌入一个实现业务逻辑的结构体（biz层）
	uc  *biz.TodoUsecase
	log *log.Helper
}

func NewTodoService(uc *biz.TodoUsecase, logger log.Logger) *TodoService {
	return &TodoService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *TodoService) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoReply, error) {
	// 请求来了
	s.log.Debugw("CeateTodo req", req)
	// 1. 请求参数的校验
	if len(req.GetTitle()) == 0 {
		return &pb.CreateTodoReply{}, errors.New("无效的title")
	}
	// 2.调用业务逻辑
	data, err := s.uc.Create(ctx, &biz.Todo{Title: req.Title})
	if err != nil {
		return nil, err
	}
	// 3.返回响应
	return &pb.CreateTodoReply{
		Id:     data.ID,
		Title:  data.Title,
		Status: data.Status,
	}, nil
}
func (s *TodoService) UpdateTodo(ctx context.Context, req *pb.UpdateTodoRequest) (*pb.UpdateTodoReply, error) {
	err := s.uc.Update(ctx, &biz.Todo{
		ID:     req.Id,
		Title:  req.Title,
		Status: req.Status,
	})
	return &pb.UpdateTodoReply{}, err
}
func (s *TodoService) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*pb.DeleteTodoReply, error) {
	err := s.uc.Delete(ctx, req.Id)
	return &pb.DeleteTodoReply{}, err
}
func (s *TodoService) GetTodo(ctx context.Context, req *pb.GetTodoRequest) (*pb.GetTodoReply, error) {
	// 1. 参数处理
	if req.Id <= 0 {
		return &pb.GetTodoReply{}, errors.New("无效的id")
	}
	// 2. 调用biz层业务逻辑
	ret, err := s.uc.Get(ctx, req.Id)
	if err != nil {
		// return nil, err
		// 返回自定义错误
		return nil, v1.ErrorTodoNotFound("id:%v todo is not found", req.Id)
	}
	// 3. 按格式返回响应
	return &pb.GetTodoReply{
		Todo: &pb.Todo{
			Id:     ret.ID,
			Title:  ret.Title,
			Status: ret.Status,
		},
	}, nil
}
func (s *TodoService) ListTodo(ctx context.Context, req *pb.ListTodoRequest) (*pb.ListTodoReply, error) {
	dataList, err := s.uc.List(ctx)
	if err != nil {
		return nil, err
	}
	reply := &pb.ListTodoReply{}
	for _, data := range dataList {
		reply.Data = append(reply.Data, &pb.Todo{
			Id:     data.ID,
			Title:  data.Title,
			Status: data.Status,
		})
	}
	return reply, nil
}
