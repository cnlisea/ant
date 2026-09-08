package mysql

import (
	"context"
	"database/sql"
	"errors"
)

type DB struct {
	ctx context.Context
	db  *sql.DB
}

func (db *DB) Query(query string, args ...any) (*DBRows, error) {
	var (
		tx  *sql.Tx
		err error
	)
	db.ctx, tx, err = db._TxBegin(db.ctx)
	if err != nil {
		return &DBRows{
			ctx: db.ctx,
		}, err
	}
	rows, err := tx.QueryContext(db.ctx, query, args...)
	return &DBRows{
		ctx:  db.ctx,
		Rows: rows,
	}, err
}

func (db *DB) QueryRow(query string, args ...any) *DBRow {
	rows, err := db.Query(query, args...)
	return &DBRow{
		ctx:  rows.Ctx(),
		err:  err,
		rows: rows,
	}
}

func (db *DB) Exec(query string, args ...any) (*DBResult, error) {
	var (
		tx  *sql.Tx
		err error
	)
	db.ctx, tx, err = db._TxBegin(db.ctx)
	if err != nil {
		return &DBResult{
			ctx:    db.ctx,
			Result: nil,
		}, err
	}

	result, err := tx.ExecContext(db.ctx, query, args...)
	return &DBResult{
		ctx:    db.ctx,
		Result: result,
	}, err
}

func (db *DB) Close() error {
	tx := db._TxCtx(db.ctx)
	if tx == nil {
		return nil
	}

	if db.TxErrVal(db.ctx) != nil {
		return db._TxRollBack(db.ctx)
	}
	return db._TxCommit(db.ctx)
}

func (db *DB) Begin() (*sql.Tx, error) {
	return nil, errors.New("Begin is not supported by this db")
}
