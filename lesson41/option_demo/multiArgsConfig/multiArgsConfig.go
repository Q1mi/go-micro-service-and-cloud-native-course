package multiArgsConfig

type MultiargsConfig struct {
	a string
	b string
	c int
}
const defaultValueC = 1

func NewServiceConfig(a, b string, c ...int) *MultiargsConfig {
	valueC := defaultValueC
	if len(c) > 0 {
		valueC = c[0]
	}
	return &MultiargsConfig{
		a: a,
		b: b,
		c: valueC,
	}
}