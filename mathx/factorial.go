package mathx

import (
	"golang.org/x/exp/constraints"
)

func Factorial[T constraints.Integer](n T) T {
	var (
		val T = 1
		i   T
	)
	for i = 2; i <= n; i++ {
		val *= i
	}
	return val
}
