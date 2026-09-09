package mysql

import (
	"database/sql"
)

type DBRows struct {
	*sql.Rows
}
