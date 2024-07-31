package grpc_server

import (
	"golang.org/x/exp/slog"

	"github.com/neracastle/auth/internal/usecases"
	userdesc "github.com/neracastle/auth/pkg/user_v1"
)

// Server GRPC сервер с ручками сервиса auth
type Server struct {
	userdesc.UnimplementedUserV1Server
	logger *slog.Logger
	srv    usecases.UserService
}

// NewServer новый экземпляр grpc сервера
func NewServer(logger *slog.Logger, srv usecases.UserService) *Server {
	return &Server{logger: logger, srv: srv}
}
