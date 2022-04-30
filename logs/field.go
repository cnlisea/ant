package logs

import (
	"time"

	"go.uber.org/zap"
)

type Field = zap.Field

func Error(key string, val error) Field {
	return zap.NamedError(key, val)
}

func Errors(key string, val []error) Field {
	return zap.Errors(key, val)
}

func Binary(key string, val []byte) Field {
	return zap.Binary(key, val)
}

func Bool(key string, val bool) Field {
	return zap.Bool(key, val)
}

func ByteString(key string, val []byte) Field {
	return zap.ByteString(key, val)
}

func Float64(key string, val float64) Field {
	return zap.Float64(key, val)
}

func Float32(key string, val float32) Field {
	return zap.Float32(key, val)
}

func Int(key string, val int) Field {
	return zap.Int(key, val)
}

func Int64(key string, val int64) Field {
	return zap.Int64(key, val)
}

func Int32(key string, val int32) Field {
	return zap.Int32(key, val)
}

func Int16(key string, val int16) Field {
	return zap.Int16(key, val)
}

func Int8(key string, val int8) Field {
	return zap.Int8(key, val)
}

func String(key string, val string) Field {
	return zap.String(key, val)
}

func Uint(key string, val uint) Field {
	return zap.Uint(key, val)
}

func Uint64(key string, val uint64) Field {
	return zap.Uint64(key, val)
}

func Uint32(key string, val uint32) Field {
	return zap.Uint32(key, val)
}

func Uint16(key string, val uint16) Field {
	return zap.Uint16(key, val)
}

func Uint8(key string, val uint8) Field {
	return zap.Uint8(key, val)
}

func Bools(key string, bs []bool) Field {
	return zap.Bools(key, bs)
}

func Ints(key string, nums []int) Field {
	return zap.Ints(key, nums)
}

func Int64s(key string, nums []int64) Field {
	return zap.Int64s(key, nums)
}

func Int32s(key string, nums []int32) Field {
	return zap.Int32s(key, nums)
}

func Int16s(key string, nums []int16) Field {
	return zap.Int16s(key, nums)
}

func Int8s(key string, nums []int8) Field {
	return zap.Int8s(key, nums)
}

func Strings(key string, ss []string) Field {
	return zap.Strings(key, ss)
}

func Times(key string, ts []time.Time) Field {
	return zap.Times(key, ts)
}

func Uints(key string, nums []uint) Field {
	return zap.Uints(key, nums)
}

func Uint64s(key string, nums []uint64) Field {
	return zap.Uint64s(key, nums)
}

func Uint32s(key string, nums []uint32) Field {
	return zap.Uint32s(key, nums)
}

func Uint16s(key string, nums []uint16) Field {
	return zap.Uint16s(key, nums)
}

func Uint8s(key string, nums []uint8) Field {
	return zap.Uint8s(key, nums)
}
