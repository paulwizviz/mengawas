package acidity

import "errors"

var (
	ErrInvalidValue = errors.New("invalid value")
)

type Unit struct {
	value float32
}

func (u Unit) Value() float32 {
	return u.value
}

func NewUnit(value float32) (Unit, error) {
	if value < 0 || value > 14 {
		return Unit{}, ErrInvalidValue
	}
	return Unit{value: value}, nil
}
