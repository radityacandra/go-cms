package database

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// combine tx and db executor
type QueryExecutor interface {
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
}

func (db *DB) BeginTransaction(ctx context.Context) (context.Context, error) {
	if ctx.Value(TX_CONTEXT_KEY) == nil {
		tx, err := db.BeginTxx(ctx, &sql.TxOptions{})
		if err != nil {
			return nil, err
		}

		txCtx := context.WithValue(ctx, TX_CONTEXT_KEY, tx)
		return txCtx, nil
	}

	return ctx, nil
}

// GetTransaction returns the transaction from context if it exists, otherwise returns nil
// This prevents creating orphaned transactions
func (db *DB) GetTransaction(ctx context.Context) QueryExecutor {
	if tx := ctx.Value(TX_CONTEXT_KEY); tx != nil {
		return tx.(*sqlx.Tx)
	}
	return db
}

func (db *DB) CommitOrRollbackTransaction(ctx context.Context, err error) error {
	if tx := db.GetTransaction(ctx).(*sqlx.Tx); tx == nil {
		return nil
	}

	tx := db.GetTransaction(ctx).(*sqlx.Tx)

	if err == nil {
		return tx.Commit()
	} else {
		return tx.Rollback()
	}
}

func (db *DB) IsTransaction(ctx context.Context) bool {
	if tx := ctx.Value(TX_CONTEXT_KEY); tx != nil {
		return true
	}

	return false
}
