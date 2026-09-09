package proxy

import (
	"context"
	"database/sql"
	"github.com/cnlisea/ant/app/db/mysql"
)

type DBMySQL interface {
	GetDB(name ...string) *sql.DB
	GetDBTx(ctx *context.Context, name ...string) *mysql.DB
}
