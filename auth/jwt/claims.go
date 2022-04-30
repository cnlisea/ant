package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Uid uint32 `json:"uid"`
	jwt.StandardClaims
}
