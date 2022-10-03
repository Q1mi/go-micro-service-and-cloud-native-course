package main

import (
	"context"
	"errors"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	defaulShelfSize = 100
)

// 使用GORM

func NewDB(dsn string) (*gorm.DB, error) {
	// mysql
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// sqlite
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Shelf{}, &Book{})
	return db, nil
}

// 定义Model

// Shelf 书架
type Shelf struct {
	ID       int64 `gorm:"primaryKey"`
	Theme    string
	Size     int64
	CreateAt time.Time
	UpdateAt time.Time
}

// Book 图书
type Book struct {
	ID       int64 `gorm:"primaryKey"`
	Author   string
	Title    string
	ShelfID  int64
	CreateAt time.Time
	UpdateAt time.Time
}

// 数据库操作
type bookstore struct {
	db *gorm.DB
}

// CreateShelf 创建书架
func (b *bookstore) CreateShelf(ctx context.Context, data Shelf) (*Shelf, error) {
	if len(data.Theme) <= 0 {
		return nil, errors.New("invalid theme")
	}
	size := data.Size
	if size <= 0 {
		size = defaulShelfSize
	}
	v := Shelf{Theme: data.Theme, Size: size, CreateAt: time.Now(), UpdateAt: time.Now()}
	err := b.db.WithContext(ctx).Create(&v).Error
	return &v, err
}

// GetShelf 获取书架
func (b *bookstore) GetShelf(ctx context.Context, id int64) (*Shelf, error) {
	v := Shelf{}
	err := b.db.WithContext(ctx).First(&v, id).Error
	return &v, err
}

// ListShelves 书架列表
func (b *bookstore) ListShelves(ctx context.Context) ([]*Shelf, error) {
	var vl []*Shelf
	err := b.db.WithContext(ctx).Find(&vl).Error
	return vl, err
}

// DeleteShelf 删除书架
func (b *bookstore) DeleteShelf(ctx context.Context, id int64) error {
	return b.db.WithContext(ctx).Delete(&Shelf{}, id).Error
}

// GetBookListByShelfID 根据书架id查询图书
func (b *bookstore) GetBookListByShelfID(ctx context.Context, shelfID int64, cursor string, pageSize int) ([]*Book, error) {
	var vl []*Book
	err := b.db.Debug().WithContext(ctx).Where("shelf_id = ? and id > ?", shelfID, cursor).Order("id asc").Limit(pageSize).Find(&vl).Error
	return vl, err
}
