package optionConfig

type optioncofig struct {
	a string
	b string
	c int

	x Info
}

type Info struct {
	Addr string
}

const defaultValueC = 1

type FuncServiceConfigOption func(*optioncofig)

func NewServiceConfig(a, b string, opts ...FuncServiceConfigOption) *optioncofig {
	sc := &optioncofig{
		a: a,
		b: b,
		c: defaultValueC,
	}

	// 针对可能传进来的FuncServiceConfigOption参数做处理
	for _, opt := range opts {
		opt(sc)
	}
	return sc
}

// 针对可选的配置实现一些专用的配置方法
func WithC(c int) FuncServiceConfigOption {
	return func(sc *optioncofig) {
		sc.c = c // 把传进来的ServiceConfig对象的C字段修改
	}
}
func WithX(info Info) FuncServiceConfigOption {
	return func(sc *optioncofig) {
		sc.x = info
	}
}
