package demo

import (
	"errors"

	"github.com/google/wire"
)

type X struct {
	Value int
}

func NewX() X {
	return X{Value: 7}
}

type Y struct {
	Value int
}

func NewY(x X) Y {
	return Y{Value: x.Value + 1}
}

type Z struct {
	Value int
}

func NewZ(y Y) (Z, error) {
	if y.Value == 0 {
		return Z{}, errors.New("bad y")
	}
	return Z{Value: y.Value + 2}, nil
}

var ProviderSet = wire.NewSet(NewX, NewY, NewZ)
