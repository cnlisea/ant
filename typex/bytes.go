package typex

import "unsafe"

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func BytesToInts(b []byte) []int {
	if b == nil {
		return nil
	}

	var (
		bLen = len(b)
		r    = make([]int, bLen)
		i    int
	)
	for i = 0; i < bLen; i++ {
		r[i] = int(b[i])
	}

	return r
}
