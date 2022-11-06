# lesson41

Go语言中的Option模式

## 解决的问题
1. Go语言中如何设置函数默认参数

```python
def foo(msg, name="q1mi"):
    pass

foo("hello")
foo("hello", name="七米")
```
2. 结构体字段会随着业务发展增加字段


Option模式
1. 利用的Go函数的不定长参数
```go
// Dial creates a client connection to the given target.
func Dial(target string, opts ...DialOption) (*ClientConn, error) {
	return DialContext(context.Background(), target, opts...)
}

Dial("consul://127.0.0.1:8500/hello")
grpc.Dial(
    "consul://localhost:8500/hello?healthy=true", // grpc中使用consul名称解析器，
    // 指定负载均衡策略，这里使用的是gRPC自带的round_robin
    grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
    grpc.WithTransportCredentials(insecure.NewCredentials()),
)

// 除了target之外的其他的参数会以 []DialOption 方式赋值给 opts
```


# 进阶版Option

```go
sc.C = 100 // 我就不用你提供的WithXxx方法，我可以直接改！ 你奈我何！！！☺️
```

我的结构体和结构体字段都是小写字母开头，不对外可见（对包外不可见）
用接口类型去隐藏具体实现

