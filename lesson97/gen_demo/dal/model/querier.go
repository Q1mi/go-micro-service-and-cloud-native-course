// dal/model/querier.go

package model

import "gorm.io/gen"

// 通过添加注释生成自定义方法
// GEN框架读取你写的注释，自动帮你生成一个注释下的方法的具体实现

type Querier interface {
	// SELECT * FROM @@table WHERE id=@id
	GetByID(id int) (gen.T, error) // 返回结构体和error

	// GetByIDReturnMap 根据ID查询返回map
	//
	// SELECT * FROM @@table WHERE id=@id
	GetByIDReturnMap(id int) (gen.M, error) // 返回 map 和 error

	// SELECT * FROM @@table WHERE author=@author
	GetBooksByAuthor(author string) ([]*gen.T, error) // 返回数据切片和 error
}
