package database

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

type ITransaction interface {
	BeginTransaction(ctx context.Context) (context.Context, error)
	GetTransaction(ctx context.Context) QueryExecutor
	CommitOrRollbackTransaction(ctx context.Context, err error) error
	IsTransaction(ctx context.Context) bool
}

type CTX_KEY string

const TX_CONTEXT_KEY CTX_KEY = "TX_KEY"

func Init(ctx context.Context, dbDsn string) (*DB, error) {
	db, err := sql.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return &DB{sqlx.NewDb(db, "pgx")}, nil
}
