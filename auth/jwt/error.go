package jwt

import (
	"errors"
)

var (
	ErrTokenExpired = errors.New("token expired")
)
