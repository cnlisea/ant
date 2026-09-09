package mysql

import (
	"database/sql"
	"errors"
)

type DBRow struct {
	err  error
	rows *DBRows
}

func (r *DBRow) Scan(dest ...any) error {
	if r.err != nil {
		return r.err
	}

	defer r.rows.Close()
	for _, dp := range dest {
		if _, ok := dp.(*sql.RawBytes); ok {
			return errors.New("sql: RawBytes isn't allowed on Row.Scan")
		}
	}

	if !r.rows.Next() {
		if err := r.rows.Err(); err != nil {
			return err
		}
		return sql.ErrNoRows
	}
	err := r.rows.Scan(dest...)
	if err != nil {
		return err
	}
	// Make sure the query can be processed to completion with no errors.
	return r.rows.Close()
}

func (r *DBRow) Err() error {
	return r.err
}
