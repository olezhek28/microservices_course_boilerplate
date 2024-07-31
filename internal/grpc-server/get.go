package grpc_server

import (
	"context"

	"github.com/neracastle/go-libs/pkg/sys/logger"
	"golang.org/x/exp/slog"

	userdesc "github.com/neracastle/auth/pkg/user_v1"
)

// Get возвращает данные клиента
func (s *Server) Get(ctx context.Context, req *userdesc.GetRequest) (*userdesc.GetResponse, error) {
	ctx = logger.AssignLogger(ctx, s.logger)

	log := s.logger.With(slog.String("method", "grpc-server.Get"))
	log.Debug("called", slog.Any("user_id", req.GetId()))

	dbUser, err := s.srv.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	rsp := FromUsecaseToGetResponse(dbUser)

	return rsp, nil
}
