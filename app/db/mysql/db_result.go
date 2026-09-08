package mysql

import (
	"database/sql"
	"golang.org/x/net/context"
)

type DBResult struct {
	ctx context.Context
	sql.Result
}

func (r *DBResult) Ctx() context.Context {
	return r.ctx
}
