package mathx

import (
	"golang.org/x/exp/constraints"
)

func PermutNum[T constraints.Integer](n, m T) T {
	return Factorial[T](n) / Factorial[T](n-m)
}
