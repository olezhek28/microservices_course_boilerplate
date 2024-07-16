package logger

import (
	"context"

	"golang.org/x/exp/slog"
)

type loggerKey struct{}

// AssignLogger прокидывает логгер в контекст
func AssignLogger(ctx context.Context, logger *slog.Logger) context.Context {
	ctx = context.WithValue(ctx, loggerKey{}, logger)

	return ctx
}

// GetLogger получает логгер из контекста
func GetLogger(ctx context.Context) *slog.Logger {
	return ctx.Value(loggerKey{}).(*slog.Logger)
}
