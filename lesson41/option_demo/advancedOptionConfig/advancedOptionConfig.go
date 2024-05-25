package advancedOptionConfig

const defaultValueName = "七米"

type AdvancedOptionConfig struct {
	name string
	age  int
}

func NewConfig(age int, opts ...ConfigOption) *AdvancedOptionConfig {
	cfg := &AdvancedOptionConfig{
		name: defaultValueName,
		age:  age,
	}

	for _, opt := range opts {
		opt.apply(cfg)
	}
	return cfg
}

type ConfigOption interface {
	apply(*AdvancedOptionConfig)
}

type funcOption struct {
	f func(*AdvancedOptionConfig)
}

func (f funcOption) apply(cfg *AdvancedOptionConfig) {
	f.f(cfg)
}

func NewfuncOption(f func(*AdvancedOptionConfig)) funcOption {
	return funcOption{f: f}
}

func WithConfigName(name string) ConfigOption {
	return NewfuncOption(func(cfg *AdvancedOptionConfig) {
		cfg.name = name
	})
}
