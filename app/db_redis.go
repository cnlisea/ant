package app

import (
	"errors"

	"github.com/cnlisea/ant/app/db/redis"
)

func (a *App) DBRedisRegister(name string,
	password string, addr string, port uint16, db int,
	active int, idle int, idleTimeout int) error {
	if a.redis != nil && a.redis[name] != nil {
		return ErrRedisInstanceAlreadyExisted
	}

	client, err := redis.NewClient(password, addr, port, db, active, idle, idleTimeout)
	if err != nil {
		return errors.New("redis client new fail, err:" + err.Error())
	}

	if a.redis == nil {
		a.redis = make(map[string]*redis.Client, 1)
	}
	a.redis[name] = client
	return nil
}

func (a *App) DBRedisInstance(name ...string) *redis.Client {
	if a.redis == nil {
		return nil
	}

	var key string
	if name != nil && len(name) > 0 {
		key = name[0]
	}

	return a.redis[key]
}

func (a *App) DBRedisAllClose() {
	ctx := a.Context()
	for _, c := range a.redis {
		c.Close(ctx)
	}
}
