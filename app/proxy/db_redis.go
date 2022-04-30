package proxy

import (
	"github.com/gomodule/redigo/redis"
)

type DBRedis interface {
	GetDB(name ...string) redis.Conn
}
