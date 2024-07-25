package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Client клиент БД
type Client interface {
	DB() DB
	Close() error
}

// Handler - функция, которая выполняется в транзакции
type Handler func(ctx context.Context) error

// TxManager менеджер транзакций
type TxManager interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

// Transactor интерфейс для работы с транзакциями
type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

// Query обертка над запросом, хранящая имя запроса и сам запрос
type Query struct {
	Name     string
	QueryRaw string
}

// QueryExecer интерфейс для работы с обычными запросами
type QueryExecer interface {
	Exec(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

// Pinger интерфейс для проверки соединения с БД
type Pinger interface {
	Ping(ctx context.Context) error
}

// DB интерфейс для работы с БД
type DB interface {
	QueryExecer
	Transactor
	TxManager
	Pinger
	Close()
}
