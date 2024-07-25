package pg

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	db "github.com/neracastle/auth/internal/client"
)

type pgClient struct {
	masterDb db.DB
}

func NewClient(ctx context.Context, dsn string) (db.Client, error) {
	connect, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	pgc := &pgClient{masterDb: NewDB(connect)}

	return pgc, nil
}

func (c *pgClient) DB() db.DB {
	return c.masterDb
}

func (c *pgClient) Close() error {
	if c.masterDb != nil {
		c.masterDb.Close()
	}

	return nil
}
