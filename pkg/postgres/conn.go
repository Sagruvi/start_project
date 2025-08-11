package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	// Optionally check if the connection is valid
	conn, err := pool.Acquire(ctx)
	if err != nil {
		pool.Close()
		return nil, err
	}
	defer conn.Release()

	return pool, nil
}
