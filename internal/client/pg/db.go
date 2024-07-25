package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/exp/slog"

	"github.com/neracastle/auth/internal/app/logger"
	db "github.com/neracastle/auth/internal/client"
)

var _ db.DB = (*pg)(nil)

type key string

const (
	TxCtxKey key = "tx"
)

type pg struct {
	pool *pgxpool.Pool
}

func NewDB(dbc *pgxpool.Pool) db.DB {
	return &pg{pool: dbc}
}

func (p *pg) Exec(ctx context.Context, q db.Query, args ...interface{}) (pgconn.CommandTag, error) {
	logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxCtxKey).(pgx.Tx)
	if ok {
		return tx.Exec(ctx, q.QueryRaw, args...)
	}

	return p.pool.Exec(ctx, q.QueryRaw, args...)
}

func (p *pg) Query(ctx context.Context, q db.Query, args ...interface{}) (pgx.Rows, error) {
	logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxCtxKey).(pgx.Tx)
	if ok {
		return tx.Query(ctx, q.QueryRaw, args...)
	}

	return p.pool.Query(ctx, q.QueryRaw, args...)
}

func (p *pg) QueryRow(ctx context.Context, q db.Query, args ...interface{}) pgx.Row {
	logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxCtxKey).(pgx.Tx)
	if ok {
		return tx.QueryRow(ctx, q.QueryRaw, args...)
	}

	return p.pool.QueryRow(ctx, q.QueryRaw, args...)
}

func (p *pg) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.pool.BeginTx(ctx, txOptions)
}

func (p *pg) Ping(ctx context.Context) error {
	return p.pool.Ping(ctx)
}

func (p *pg) Close() {
	p.pool.Close()
}

func MakeContextTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TxCtxKey, tx)
}

func logQuery(ctx context.Context, q db.Query, args ...interface{}) {
	log := logger.GetLogger(ctx)
	log.Debug("db query",
		slog.String("name", q.Name),
		slog.String("query", q.QueryRaw),
		slog.String("args", fmt.Sprintln(args...)))
}
