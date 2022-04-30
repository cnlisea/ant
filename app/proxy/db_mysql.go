package proxy

import "database/sql"

type DBMySQL interface {
	GetDB(name ...string) *sql.DB
}
