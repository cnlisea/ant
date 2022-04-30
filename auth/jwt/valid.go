package jwt

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func (j *Jwt) Valid(tokenStr string) (uint32, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return j.Key, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok && ve.Errors&jwt.ValidationErrorExpired != 0 {
			return 0, ErrTokenExpired
		}
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("token invalid")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return 0, errors.New("token claims type invalid")
	}

	return claims.Uid, nil
}
