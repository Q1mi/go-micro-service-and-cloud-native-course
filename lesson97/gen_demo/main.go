package main

import (
	"context"
	"fmt"
	"gen-demo/dal/model"
	"gen-demo/dal/query"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const MySQLDSN = "root:root1234@tcp(127.0.0.1:13306)/db2?charset=utf8mb4&parseTime=True"

func connectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	return db
}

func main() {
	fmt.Println("gen demo start...")
	db := connectDB(MySQLDSN)

	// 重要！
	query.SetDefault(db)

	// CRUD
	// 新增
	// 结构体
	// b1 := model.Book{
	// 	Title:       "《Go语言之路》",
	// 	Author:      "七米",
	// 	Price:       100,
	// 	PublishDate: time.Now(),
	// }
	// err := query.Book.
	// 	WithContext(context.Background()).
	// 	Create(&b1)
	// if err != nil {
	// 	fmt.Printf("create book fail, err:%v\n", err)
	// 	return
	// }
	// fmt.Printf("b1:%#v\n", b1)

	// 查询
	// 方法1
	b, err := query.Book.
		WithContext(context.Background()).
		Where(query.Book.ID.Eq(7)).
		First()
	fmt.Printf("b:%#v err:%v\n", b, err)
	// 方法2
	book := query.Book
	b, err = book.
		WithContext(context.Background()).
		Where(book.ID.Eq(7)).
		First()
	fmt.Printf("b:%#v err:%v\n", b, err)

	// 更新
	ret, err := query.Book.
		WithContext(context.Background()).
		Where(query.Book.ID.Eq(7)).
		Update(query.Book.Price, 200)
	if err != nil {
		fmt.Printf("update book fail, err:%v\n", err)
		return
	}
	fmt.Println(ret.RowsAffected)

	// 删除
	// 方法1
	b3 := model.Book{ID: 3}
	ret, err = query.Book.
		WithContext(context.Background()).
		Delete(&b3)
	if err != nil {
		fmt.Printf("delete book fail, err:%v\n", err)
		return
	}
	fmt.Println(ret.RowsAffected)
	// 方法2
	ret, err = query.Book.
		WithContext(context.Background()).
		Where(query.Book.ID.Eq(7)).
		Delete()
	if err != nil {
		fmt.Printf("delete book fail, err:%v\n", err)
		return
	}
	fmt.Println(ret.RowsAffected)
}
