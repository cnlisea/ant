package mathx

import (
	"golang.org/x/exp/constraints"
)

func CombinNum[T constraints.Integer](n, m T) T {
	return Factorial[T](n) / (Factorial[T](n-m) * Factorial[T](m))
}

func Combin[T constraints.Integer](data []T, n int) [][]T {
	if n <= 0 {
		return nil
	}

	dataLen := len(data)
	if dataLen == 0 || dataLen < n {
		return nil
	}

	var (
		num   = CombinNum[int](dataLen, n)
		index = make([][]uint8, num)
		i, j  int
	)
	index[0] = make([]uint8, dataLen)
	for i = 0; i < n; i++ {
		index[0][i] = 1
	}
	for i = 1; i < num; i++ {
		if index[i] == nil {
			index[i] = make([]uint8, dataLen)
			copy(index[i], index[i-1])
		}
		for j = 0; j < dataLen-1; j++ {
			if index[i][j] == 1 && index[i][j+1] == 0 {
				index[i][j], index[i][j+1] = index[i][j+1], index[i][j]
				break
			}
		}
	}

	var (
		vals = make([][]T, num)
		val  []T
	)
	for i = 0; i < num; i++ {
		val = make([]T, 0, n)
		for j = 0; j < dataLen; j++ {
			if index[i][j] == 1 {
				val = append(val, data[j])
			}
		}
		vals[i] = val
	}
	return vals
}
