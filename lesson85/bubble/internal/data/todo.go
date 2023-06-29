package data

import (
	"context"
	"fmt"

	"bubble/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type todoRepo struct {
	data *Data
	log  *log.Helper
}

// NewTodoRepo .
func NewTodoRepo(data *Data, logger log.Logger) biz.TodoRepo {
	return &todoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *todoRepo) Save(ctx context.Context, t *biz.Todo) (*biz.Todo, error) {
	// 实现数据库的操作
	fmt.Printf("save: t:%#v\n", t)
	return t, nil
}

func (r *todoRepo) Update(ctx context.Context, t *biz.Todo) error {
	return nil
}

func (r *todoRepo) FindByID(context.Context, int64) (*biz.Todo, error) {
	return nil, nil
}

func (r *todoRepo) Delete(context.Context, int64) error {
	return nil
}

func (r *todoRepo) ListAll(context.Context) ([]*biz.Todo, error) {
	return nil, nil
}
