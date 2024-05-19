package basicConfig

type BasicConfig struct {
	a string
	b string
	c int
}

// NewServiceConfig 创建一个ServiceConfig的函数
func NewServiceConfig(a, b string, c int) *BasicConfig {
	return &BasicConfig{
		a: a,
		b: b,
		c: c,
	}
}