package typex

import (
	"encoding/hex"
	"strconv"
	"strings"
)

func Int64ToByte(data int64) ([]byte, error) {
	var (
		hexString = strconv.FormatInt(data, 16)
		lenString = len(hexString)
	)
	if lenString > 8 {
		return nil, ErrInvalidData
	}

	if lenString < 8 {
		hexString = strings.Repeat("0", 8-lenString) + hexString
	}

	return hex.DecodeString(hexString)
}
