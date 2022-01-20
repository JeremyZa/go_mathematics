package operations

import (
	"errors"
)

func addition(lhs int, rhs int) int {
	return lhs + rhs
}

func substraction(lhs int, rhs int) int {
	return lhs - rhs
}

func multiplication(lhs int, rhs int) int {
	return lhs * rhs
}

func division(lhs float64, rhs float64) (float64, error) {
	if rhs == 0 {
		return 0, errors.New("can't divide by 0")
	}
	return lhs / rhs, nil
}
