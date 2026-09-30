package typex

import "time"

func TimeCurTsUint32() uint32 {
	return uint32(time.Now().Unix())
}
