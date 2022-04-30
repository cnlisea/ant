package app

import (
	"errors"

	"github.com/cnlisea/ant/app/db/mongo"
)

func (a *App) DBMongoRegister(name string,
	user string, password string,
	addr []string, dbName string,
	replicaSet string, connTimeout int,
	active int, idle int, idleTimeout int) error {
	if a.mongo != nil && a.mongo[name] != nil {
		return ErrMongoInstanceAlreadyExisted
	}
	client, err := mongo.NewClient(a.ctx, user, password, addr, dbName, replicaSet, connTimeout, active, idle, idleTimeout)
	if err != nil {
		return errors.New("mongo client new fail, err:" + err.Error())
	}

	if a.mongo == nil {
		a.mongo = make(map[string]*mongo.Client, 1)
	}
	a.mongo[name] = client
	return nil
}

func (a *App) DBMongoInstance(name ...string) *mongo.Client {
	if a.mongo == nil {
		return nil
	}

	if name != nil && len(name) > 0 {
		return a.mongo[name[0]]
	}

	return a.mongo[""]
}

func (a *App) DBMongoAllClose() {
	ctx := a.Context()
	for _, c := range a.mongo {
		c.Close(ctx)
	}
}
