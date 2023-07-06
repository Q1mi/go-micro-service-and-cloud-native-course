//go:build wireinject
// +build wireinject

package main

import (
	"wire_demo/demo"

	"github.com/google/wire"
)

// 注入器函数
func initZ() (demo.Z, error) {
	/*
		x := demo.NewX()
		y := demo.NewY(x)
		z, err := demo.NewZ(y)
	*/
	// wire.Build(demo.NewX, demo.NewY, demo.NewZ)
	panic(wire.Build(demo.ProviderSet))
	// return demo.Z{}, nil
}
