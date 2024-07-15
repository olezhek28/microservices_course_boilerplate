package logger

import (
	"context"

	"golang.org/x/exp/slog"
)

func AssignLogger(ctx context.Context, logger *slog.Logger) context.Context {
	ctx = context.WithValue(ctx, "logger", logger)

	return ctx
}

func GetLogger(ctx context.Context) *slog.Logger {

	v := ctx.Value("logger")

	if v != nil {
		return v.(*slog.Logger)
	}

	return ctx.Value("logger").(*slog.Logger)
}
