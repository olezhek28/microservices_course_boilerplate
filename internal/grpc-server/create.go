package grpc_server

import (
	"context"
	"errors"

	"golang.org/x/exp/slog"

	"github.com/neracastle/auth/internal/app/logger"
	userdesc "github.com/neracastle/auth/pkg/user_v1"
)

// Create регистрирует нового пользователя
func (s *Server) Create(ctx context.Context, req *userdesc.CreateRequest) (*userdesc.CreateResponse, error) {
	ctx = logger.AssignLogger(ctx, s.logger)

	log := s.logger.With(slog.String("method", "grpc-server.Create"))
	log.Debug("called", slog.Any("req", req))

	if req.GetRole() == userdesc.Role_UNKNOWN {
		return nil, errors.New("роль не задана")
	}

	apiDto := FromGrpcToCreateUsecase(req)
	userId, err := s.srv.Create(ctx, apiDto)

	if err != nil {
		return nil, err
	}

	return &userdesc.CreateResponse{Id: userId}, nil
}
