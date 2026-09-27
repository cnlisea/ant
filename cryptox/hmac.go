package cryptox

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"

	"github.com/cnlisea/ant/typex"
)

func HmacSha256(data string, key string) string {
	h := hmac.New(sha256.New, typex.StringToBytes(key))

	h.Write(typex.StringToBytes(data))
	h.Write(typex.StringToBytes(key))

	return fmt.Sprintf("%x", h.Sum(nil))
}
