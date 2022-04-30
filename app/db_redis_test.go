package app

import (
	"testing"

	"github.com/cnlisea/ant/logs"
	"github.com/gomodule/redigo/redis"
)

func TestApp_DBRedisRegister(t *testing.T) {
	var (
		app = New()
		err error
	)
	if err = app.Logger("stdout", logs.LevelDebug, true, 0); err != nil {
		t.Fatal(err)
	}

	if err = app.DBRedisRegister("", "", "", 6379, 0, 10, 10, 600); err != nil {
		t.Fatal("redis register fail", err)
	}

	proxy := app.ProxyDBRedis()
	db := proxy.GetDB("")
	if db == nil {
		t.Fatal("db not found")
	}

	info, err := redis.String(db.Do("info"))
	if err != nil {
		t.Fatal("redis info fail", err)
	}

	t.Log(info)
}
