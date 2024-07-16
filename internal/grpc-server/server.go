package grpc_server

import (
	"github.com/jackc/pgx/v5"
	"golang.org/x/exp/slog"

	userdesc "github.com/neracastle/auth/pkg/user_v1"
)

// Server GRPC сервер с ручками сервиса auth
type Server struct {
	userdesc.UnimplementedUserV1Server
	logger *slog.Logger
	pgcon  *pgx.Conn
}

func (s *Server) GetLogger() *slog.Logger {
	return s.logger
}

func NewServer(logger *slog.Logger, conn *pgx.Conn) *Server {
	return &Server{logger: logger, pgcon: conn}
}
