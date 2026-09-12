package http

import (
	"net/http"
	"strings"
)

func BearerAuthToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	s := strings.SplitN(authHeader, " ", 2)
	if len(s) != 2 || strings.ToLower(s[0]) != "bearer" {
		return ""
	}

	return s[1]
}
