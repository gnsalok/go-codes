package errors

import (
	"errors"
)

var (
	ErrDivideByZero = errors.New("Divide by zero")
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}
