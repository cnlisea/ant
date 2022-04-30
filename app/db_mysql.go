package app

import (
	"errors"

	"github.com/cnlisea/ant/app/db/mysql"
)

func (a *App) DBMySqlRegister(name string,
	user string, password string,
	addr string, port uint16, dbName string,
	charset string, connTimeout string, parseTime bool, loc string,
	active int, idle int, idleTimeout int) error {
	if a.mysql != nil && a.mysql[name] != nil {
		return ErrMysqlInstanceAlreadyExisted
	}

	db, err := mysql.NewClient(user, password, addr, port, dbName, charset, connTimeout, parseTime, loc, active, idle, idleTimeout)
	if err != nil {
		return errors.New("mysql client new fail, err:" + err.Error())
	}
	if a.mysql == nil {
		a.mysql = make(map[string]*mysql.Client, 1)
	}
	a.mysql[name] = db
	return nil
}

func (a *App) DBMySqlInstance(name ...string) *mysql.Client {
	if a.mysql == nil {
		return nil
	}

	if name != nil && len(name) > 0 {
		return a.mysql[name[0]]
	}

	return a.mysql[""]
}

func (a *App) DBMySqlAllClose() {
	ctx := a.Context()
	for _, c := range a.mysql {
		c.Close(ctx)
	}
}
