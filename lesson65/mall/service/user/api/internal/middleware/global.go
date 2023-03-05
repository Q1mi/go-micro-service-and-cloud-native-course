package middleware

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

// 定义全局中间件

// 功能：
// 记录所有请求的响应信息

// rest.Middleware -> type Middleware func(next http.HandlerFunc) http.HandlerFunc
// type HandlerFunc func(ResponseWriter, *Request)

// bodyCopy 是一个自定义的结构体
// 满足 http.ResponseWriter 接口类型
type bodyCopy struct {
	http.ResponseWriter               // 结构体嵌入接口类型
	body                *bytes.Buffer // 我的小本本，用来记录响应体内容
}

func NewBodyCopy(w http.ResponseWriter) *bodyCopy {
	return &bodyCopy{
		ResponseWriter: w,
		body:           bytes.NewBuffer([]byte{}),
	}
}

func (bc bodyCopy) Write(b []byte) (int, error) {
	// 1. 先在我的小本本记录响应内容
	bc.body.Write(b)
	// 2. 再往HTTP响应里写响应内容
	return bc.ResponseWriter.Write(b)
}

// CopyResp 复制请求的响应体
func CopyResp(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 处理请求前
		// var bc = bodyCopy{ResponseWriter: w}
		// bc.Header()
		// bc.Write()
		// bc.WriteHeader()

		// 初始化得到一个自定义的 http.ResponseWriter
		bc := NewBodyCopy(w)
		next(bc, r) // 实际的路由处理handler函数
		// 处理请求后
		fmt.Printf("--> req:%v resp: %v\n", r.URL, bc.body.String())
	}
}

func MiddlewareWithAnotherService(ok bool) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if ok {
				fmt.Println("ok!")
			}
			next(w, r)
		}
	}
}
