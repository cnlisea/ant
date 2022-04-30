package ws

import "errors"

var (
	ErrConnectionClosed = errors.New("connection is closed")
	ErrClientConnClosed = errors.New("client connection is closed")
)
