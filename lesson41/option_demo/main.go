package main

import "fmt"

// option 模式

const defaultValueC = 1

// ServiceConfig 我们定义一个服务配置结构体，里面有各种配置项
type ServiceConfig struct {
	A string
	B string
	C int

	X struct{}
	Y Info
}

type Info struct {
	addr string
}

// NewServiceConfig 创建一个ServiceConfig的函数
func NewServiceConfig(a, b string, c int) *ServiceConfig {
	return &ServiceConfig{
		A: a,
		B: b,
		C: c,
	}
}

// 目标：想要A和B必须传，C爱传不传，不传就用默认值

func NewServiceConfig2(a, b string, c ...int) *ServiceConfig {
	valueC := defaultValueC
	if len(c) > 0 {
		valueC = c[0]
	}
	return &ServiceConfig{
		A: a,
		B: b,
		C: valueC,
	}
}

// Option模式

type FuncServiceConfigOption func(*ServiceConfig)

func NewServiceConfig3(a, b string, opts ...FuncServiceConfigOption) *ServiceConfig {
	sc := &ServiceConfig{
		A: a,
		B: b,
		C: defaultValueC,
	}

	// 针对可能传进来的FuncServiceConfigOption参数做处理
	for _, opt := range opts {
		opt(sc)
	}
	return sc
}

// 针对可选的配置实现一些专用的配置方法
func WithC(c int) FuncServiceConfigOption {
	return func(sc *ServiceConfig) {
		sc.C = c // 把传进来的ServiceConfig对象的C字段修改
	}
}

func WithY(info Info) FuncServiceConfigOption {
	return func(sc *ServiceConfig) {
		sc.Y = info
	}
}

func main() {
	// sc := NewServiceConfig3("qimi", "西二旗")
	// fmt.Printf("sc:%#v\n", sc)

	// 使用 WithC 函数，携带指定的c
	// sc := NewServiceConfig3("qimi", "西二旗", WithC(10))
	// fmt.Printf("sc:%#v\n", sc)

	info := Info{addr: "127.0.0.1:8080"}
	sc := NewServiceConfig3("qimi", "西二旗", WithC(10), WithY(info))
	fmt.Printf("sc:%#v\n", sc)

	sc.C = 100 // 可以直接改？？？-->

	// 进阶版Option使用
	cfg := NewConfig(18)
	fmt.Printf("cfg:%#v\n", cfg)
	cfg2 := NewConfig(18, WithConfigName("张三"))
	fmt.Printf("cfg:%#v\n", cfg2)

	// cfg2.age = ??? 在其他包中没有办法修改
}

// 进阶版Option
const defaultValueName = "七米"

type config struct {
	name string
	age  int
}

func NewConfig(age int, opts ...ConfigOption) *config {
	cfg := &config{
		name: defaultValueName,
		age:  age,
	}

	for _, opt := range opts {
		opt.apply(cfg)
	}
	return cfg
}

type ConfigOption interface {
	apply(*config)
}

type funcOption struct {
	f func(*config)
}

func (f funcOption) apply(cfg *config) {
	f.f(cfg)
}

func NewfuncOption(f func(*config)) funcOption {
	return funcOption{f: f}
}

func WithConfigName(name string) ConfigOption {
	return NewfuncOption(func(cfg *config) {
		cfg.name = name
	})
}
