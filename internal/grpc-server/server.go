package grpc_server

import (
	"github.com/neracastle/auth/internal/usecases"
	userdesc "github.com/neracastle/auth/pkg/user_v1"
)

// Server GRPC сервер с ручками сервиса auth
type Server struct {
	userdesc.UnimplementedUserV1Server
	srv usecases.UserService
}

// NewServer новый экземпляр grpc сервера
func NewServer(srv usecases.UserService) *Server {
	return &Server{srv: srv}
}
