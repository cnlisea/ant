package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (j *Jwt) Token(uid uint32, expireTime time.Duration) (string, error) {
	curTime := time.Now()
	claims := &Claims{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			IssuedAt: curTime.Unix(),
		},
	}
	if expireTime > 0 {
		claims.ExpiresAt = curTime.Add(expireTime).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(j.Key)
}
