package main

import (
	"fmt"
	"protobuf_demo/api"

	"github.com/golang/protobuf/protoc-gen-go/generator"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	fieldmask_utils "github.com/mennanov/fieldmask-utils"
)

type Book struct {
	Price int64 // ?区分默认值和0
	// Price sql.NullInt64 // 自定义结构体
	// Price *int64        // 指针
}

func foo() {
	// var book Book
	// if book.Price == nil {
	// 	// 没有赋值
	// } else {
	// 	// 赋值
	// }
	// book.Price == 0
	// book = Book{Price: 0}
	// book.Price == 0
}

// oneofDemooneof 示例
func oneofDemo() {

	// client
	// req := &api.NoticeReaderRequest{
	// 	Msg: "李文周的博客更新啦",
	// 	NoticeWay: &api.NoticeReaderRequest_Email{
	// 		Email: "123@xxx.vom",
	// 	},
	// }

	req := &api.NoticeReaderRequest{
		Msg: "李文周的博客更新啦",
		NoticeWay: &api.NoticeReaderRequest_Phone{
			Phone: "123456789",
		},
	}

	// server
	// req.NoticeWay  ?？？
	// 类型断言
	switch v := req.NoticeWay.(type) {
	case *api.NoticeReaderRequest_Email:
		noticeWithEmail(v) // 发邮件通知
	case *api.NoticeReaderRequest_Phone:
		noticeWithPhone(v)
	}
}

// wrapValueDemo 使用google/protobuf/wrappers.proto
// func wrapValueDemo() {
// 	// client
// 	book := api.Book{
// 		Title: "《跟七米学Go语言》",
// 		// Price: &wrapperspb.Int64Value{Value: 9900},
// 		Memo: &wrapperspb.StringValue{Value: "学就完事了"},
// 	}
// 	// server
// 	if book.GetPrice() == nil { // 没有给price赋值
// 		fmt.Println("没有设置price")
// 	} else {
// 		// 赋值了放心大胆的去用
// 		fmt.Println(book.GetPrice().GetValue())
// 	}
// 	if book.GetMemo() != nil {
// 		// 有设置memo
// 		fmt.Println(book.GetMemo().GetValue())
// 	}
// }

func optionalDemo() {
	// client
	book := api.Book{
		Title: "《跟七米学Go语言》",
		Price: proto.Int64(9900), // sql.NullInt64;*int64
	}
	// server
	// 如何判断book.Price 有没有被赋值呢？
	if book.Price == nil { // 没有赋值
		fmt.Println("no price")
	} else {
		fmt.Printf("book with price:%v\n", book.GetPrice())
	}
}

// fieldMaskDemo 使用field_mask实现部分更新实例
func fieldMaskDemo() {
	// client
	paths := []string{"price", "info.b", "author"} // 更新的字段信息
	req := &api.UpdateBookRequest{
		Op: "q1mi",
		Book: &api.Book{
			Author: "七米2号",
			Price:  &wrapperspb.Int64Value{Value: 8800},
			Info: &api.Book_Info{
				B: "bbbbb",
			},
		},
		UpdateMask: &fieldmaskpb.FieldMask{Paths: paths}, // 提供情报（哪些字段更新了）
	}

	// Server
	mask, _ := fieldmask_utils.MaskFromProtoFieldMask(req.UpdateMask, generator.CamelCase)
	var bookDst = make(map[string]interface{})
	// 将数据读取到map[string]interface{}
	// fieldmask-utils支持读取到结构体等，更多用法可查看文档。
	fieldmask_utils.StructToMap(mask, req.Book, bookDst)
	fmt.Printf("bookDst:%#v\n", bookDst)
}

// 发送通知相关的功能函数
func noticeWithEmail(in *api.NoticeReaderRequest_Email) {
	fmt.Printf("notice reader by email:%v\n", in.Email)
}

func noticeWithPhone(in *api.NoticeReaderRequest_Phone) {
	fmt.Printf("notice reader by phone:%v\n", in.Phone)
}

func main() {
	oneofDemo()
	wrapValueDemo()
	fieldMaskDemo()
}
