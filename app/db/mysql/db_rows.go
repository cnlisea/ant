package mysql

import (
	"database/sql"
	"golang.org/x/net/context"
)

type DBRows struct {
	ctx context.Context
	*sql.Rows
}

func (r *DBRows) Ctx() context.Context {
	return r.ctx
}
