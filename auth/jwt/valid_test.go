package jwt

import (
	"testing"
	"time"
)

func TestJwt_Valid(t *testing.T) {
	jwt := New([]byte("123456"))

	token, err := jwt.Token(1, time.Second)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("token:", token)

	t.Log(jwt.Valid(token))

	time.Sleep(2 * time.Second)

	t.Log(jwt.Valid(token))
}
