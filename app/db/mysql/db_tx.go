package mysql

import (
	"context"
	"database/sql"
)

const (
	DBTxCtxKey    = "db_tx"
	DBTxCtxErrKey = "db_tx_err"
)

func (db *DB) _TxBegin(ctx context.Context) (context.Context, *sql.Tx, error) {
	var tx *sql.Tx
	if txValue := ctx.Value(DBTxCtxKey); txValue != nil {
		tx, _ = txValue.(*sql.Tx)
	}

	var err error
	if tx == nil {
		tx, err = db.db.Begin()
		if err != nil {
			return ctx, nil, err
		}
		ctx = context.WithValue(ctx, DBTxCtxKey, tx)
	}

	return ctx, tx, nil
}

func (db *DB) TxErr(ctx context.Context, err any) context.Context {
	var tx *sql.Tx
	if txValue := ctx.Value(DBTxCtxKey); txValue != nil {
		tx, _ = txValue.(*sql.Tx)
	}

	if tx == nil {
		return ctx
	}

	return context.WithValue(ctx, DBTxCtxErrKey, err)

}

func (db *DB) TxErrVal(ctx context.Context) any {
	return ctx.Value(DBTxCtxErrKey)
}

func (db *DB) _TxCommit(ctx context.Context) error {
	tx := db._TxCtx(ctx)
	if tx == nil {
		return nil
	}

	return tx.Commit()
}

func (db *DB) _TxRollBack(ctx context.Context) error {
	tx := db._TxCtx(ctx)
	if tx == nil {
		return nil
	}

	return tx.Rollback()
}

func (db *DB) _TxCtx(ctx context.Context) *sql.Tx {
	var tx *sql.Tx
	if txValue := ctx.Value(DBTxCtxKey); txValue != nil {
		tx, _ = txValue.(*sql.Tx)
	}
	return tx
}
