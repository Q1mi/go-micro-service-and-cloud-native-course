package biz

import (
	"context"

	v1 "bubble/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Todo is a todo model.
type Todo struct {
	ID     int64
	Title  string
	Status bool
}

// TodoRepo is a todo repo.
// biz层对数据操作层提出了以下要求，不管实际的存储是MySQL还是Redis。。。
type TodoRepo interface {
	Save(context.Context, *Todo) (*Todo, error)
	Update(context.Context, *Todo) error
	Delete(context.Context, int64) error
	FindByID(context.Context, int64) (*Todo, error)
	ListAll(context.Context) ([]*Todo, error)
}

// TodoUsecase is a todo usecase.
type TodoUsecase struct {
	repo TodoRepo
	log  *log.Helper
}

// NewTodoUsecase new a todo usecase.
func NewTodoUsecase(repo TodoRepo, logger log.Logger) *TodoUsecase {
	return &TodoUsecase{repo: repo, log: log.NewHelper(logger)}
}

// Create creates a todo, and returns the new todo.
// 对外提供的业务函数
func (uc *TodoUsecase) Create(ctx context.Context, t *Todo) (*Todo, error) {
	uc.log.WithContext(ctx).Infof("Create: %#v", t)
	return uc.repo.Save(ctx, t) // 调用下一层的save方法
}
