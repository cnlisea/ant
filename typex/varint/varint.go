package varint

import (
	"math/bits"
)

func Encode(x []uint32) []byte {
	if x == nil {
		return nil
	}

	var (
		bitLen int
		bufLen int
		i      int
	)
	for i = range x {
		bitLen = bits.Len32(x[i])
		bufLen = bufLen + bitLen/7
		if bitLen%7 != 0 || bitLen == 0 {
			bufLen++
		}
	}

	var (
		buf = make([]byte, bufLen)
		v   uint32
		n   int
	)
	for i = range x {
		v = x[i]
		for v > 127 {
			buf[n] = 0x80 | uint8(v&0x7F)
			v >>= 7
			n++
		}
		buf[n] = uint8(v)
		n++
	}

	return buf
}

func Decode(b []byte) []uint32 {
	var (
		n    int
		bLen = len(b)
		i    int
	)
	for i = 0; i < bLen; i++ {
		if b[i]&0x80 == 0 {
			n++
		}
	}

	var (
		xs = make([]uint32, n)
		xi int
	)
	n = 0
	for i = 0; i < bLen; i++ {
		xs[n] = xs[n] + uint32(b[i]&0x7F)<<(7*xi)
		xi++
		if b[i]&0x80 == 0 {
			xi = 0
			n++
		}
	}
	return xs
}

func VarIntEncode(x uint32) []byte {
	if x == 0 {
		return []byte{0}
	}

	var (
		bitLen = bits.Len32(x)
		bufLen = bitLen / 7
	)
	if bitLen%7 != 0 {
		bufLen++
	}

	var (
		buf = make([]byte, bufLen)
		n   int
	)
	for x > 127 {
		buf[n] = 0x80 | uint8(x&0x7F)
		x >>= 7
		n++
	}

	buf[n] = uint8(x)
	return buf
}

func VarIntDecode(b []byte) uint32 {
	var (
		x    uint32
		bLen = len(b)
		i    int
	)
	for i = 0; i < bLen; i++ {
		x = x + uint32(b[i]&0x7F)<<(7*i)
	}
	return x
}
