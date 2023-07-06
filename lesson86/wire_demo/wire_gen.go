// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"wire_demo/demo"
)

// Injectors from wire.go:

// 注入器函数
func initZ() (demo.Z, error) {
	x := demo.NewX()
	y := demo.NewY(x)
	z, err := demo.NewZ(y)
	if err != nil {
		return demo.Z{}, err
	}
	return z, nil
}