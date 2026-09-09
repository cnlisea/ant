package mysql

import (
	"context"
	"database/sql"
	"errors"
)

type DB struct {
	ctx   *context.Context
	db    *sql.DB
	begin bool
}

func (db *DB) Query(query string, args ...any) (*DBRows, error) {
	var (
		ctx context.Context
		tx  *sql.Tx
		err error
	)
	ctx, tx, err = db._TxBegin(*(db.ctx))
	if err != nil {
		return nil, err
	}
	*(db.ctx) = ctx

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &DBRows{
		Rows: rows,
	}, nil
}

func (db *DB) QueryRow(query string, args ...any) *DBRow {
	rows, err := db.Query(query, args...)
	return &DBRow{
		err:  err,
		rows: rows,
	}
}

func (db *DB) Exec(query string, args ...any) (*DBResult, error) {
	var (
		ctx context.Context
		tx  *sql.Tx
		err error
	)
	ctx, tx, err = db._TxBegin(*(db.ctx))
	if err != nil {
		return nil, err
	}
	*(db.ctx) = ctx

	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return &DBResult{
		Result: result,
	}, nil
}

func (db *DB) Close(err *error) error {
	if !db.begin {
		return nil
	}

	ctx := *(db.ctx)
	tx := db._TxCtx(ctx)
	if tx == nil {
		return nil
	}

	if err != nil && *err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}

func (db *DB) Begin() (*sql.Tx, error) {
	return nil, errors.New("Begin is not supported by this db")
}
