package interceptors

import (
	"context"

	"github.com/neracastle/go-libs/pkg/sys/logger"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

var lg *slog.Logger

// NewLoggerInterceptor навешивает логгер на контекст запросов
func NewLoggerInterceptor(l *slog.Logger) grpc.UnaryServerInterceptor {
	lg = l
	return loggerInterceptor
}

func loggerInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log := lg
	reqID := RequestIDFromContext(ctx)
	if reqID != "" {
		log = lg.With(slog.String("request_id", reqID))
	}

	ctx = logger.AssignLogger(ctx, log)

	return handler(ctx, req)
}
