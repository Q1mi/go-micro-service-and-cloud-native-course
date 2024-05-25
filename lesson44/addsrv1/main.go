package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// go-kit addService demo 1

// 1. service
// 1.1 业务逻辑抽象为接口
type AddService interface {
	Sum(ctx context.Context, a, b int) (int, error)
}
// 1.2 实现接口
type addService struct{}
func (addService) Sum(_ context.Context, a, b int) (int, error) {
	return a + b, nil
}
// 1.3 请求和响应
type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}
type SumResponse struct {
	V   int    `json:"v"`
	Err string `json:"err,omitempty"`
}

// 2. endpoint
// 借助 适配器 将 方法 -> endpoint
func makeSumEndpoint(srv AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SumRequest)
		v, err := srv.Sum(ctx, req.A, req.B) // 方法调用
		if err != nil {
			return SumResponse{V: v, Err: err.Error()}, nil
		}
		return SumResponse{V: v}, nil
	}
}

// 3. transport
// 3.1 decode: 请求来了之后根据 协议(HTTP、HTTP2)和编码(JSON、pb、thrift) 去解析数据
func decodeSumRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request SumRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
// 3.2 encode: 把响应数据 按协议和编码 返回
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {// w: 代表响应的网络句柄, response: service层返回的响应数据
	return json.NewEncoder(w).Encode(response)
}

func main() {
	srv := addService{}

	sumHandler := httptransport.NewServer(
		makeSumEndpoint(srv),
		decodeSumRequest,
		encodeResponse,
	)

	http.Handle("/sum", sumHandler)
	http.ListenAndServe(":8080", nil)
}
