package sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SQLStore struct {
	pool    *pgxpool.Pool
	Queries *Queries
}

func NewStore(pool *pgxpool.Pool) *SQLStore {

	return &SQLStore{
		pool:    pool,
		Queries: New(pool),
	}

}

func (store *SQLStore) ExecTx(ctx context.Context, fn func(*Queries) (interface{}, error)) (interface{}, error) {

	txOptions := pgx.TxOptions{
		IsoLevel:       "repeatable read",
		AccessMode:     "read write",
		DeferrableMode: "not deferrable",
	}

	tx, err := store.pool.BeginTx(context.Background(), txOptions)

	if err != nil {
		return *new(interface{}), fmt.Errorf("failed to start transaction: %v", err)
	}

	q := New(tx)

	result, tx_err := fn(q)

	if tx_err != nil {
		tx.Rollback(context.Background())
	}

	return result, tx.Commit(context.Background())

}
