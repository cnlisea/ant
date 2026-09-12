package net

import "strconv"

func Addr(ip string, port uint16) string {
	return ip + ":" + strconv.FormatUint(uint64(port), 10)
}
