package interceptors

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type validate interface {
	Validate() error
}

// RequestIDKey ключ контекста для ID запроса
type RequestIDKey struct{}

// RequestIDInterceptor навыешивает id запроса в контекст
func RequestIDInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx = context.WithValue(ctx, RequestIDKey{}, uuid.NewString())

	if val, ok := req.(validate); ok {
		if err := val.Validate(); err != nil {
			return nil, err
		}
	}

	return handler(ctx, req)
}

// RequestIDFromContext возвращает ID запроса из контекста
func RequestIDFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	if reqID, ok := ctx.Value(RequestIDKey{}).(string); ok {
		return reqID
	}

	return ""
}
