package data

import (
	"context"

	"bubble/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// todoRepo 实现了biz层定义的repo接口
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
	// gorm操作
	err := r.data.db.Create(t).Error
	return t, err
}

func (r *todoRepo) Update(ctx context.Context, t *biz.Todo) error {
	return r.data.db.
		WithContext(ctx).
		Model(t).
		Update("status", t.Status).Error
}

func (r *todoRepo) FindByID(ctx context.Context, id int64) (*biz.Todo, error) {
	t := biz.Todo{ID: id}
	err := r.data.db.
		WithContext(ctx).
		First(&t).Error
	return &t, err
}

func (r *todoRepo) Delete(ctx context.Context, id int64) error {
	t := biz.Todo{ID: id}
	return r.data.db.
		WithContext(ctx).
		Delete(&t).Error
}

func (r *todoRepo) ListAll(ctx context.Context) ([]*biz.Todo, error) {
	var todoList []*biz.Todo
	err := r.data.db.
		WithContext(ctx).
		Find(&todoList).Error
	return todoList, err
}
