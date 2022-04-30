package typex

import (
	"encoding/hex"
	"strconv"
)

func ByteToInt64(data []byte) (int64, error) {
	return strconv.ParseInt(hex.EncodeToString(data), 16, 64)
}
