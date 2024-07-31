package grpc_server

import (
	"context"

	"github.com/neracastle/go-libs/pkg/sys/logger"
	"golang.org/x/exp/slog"
	"google.golang.org/protobuf/types/known/emptypb"

	userdesc "github.com/neracastle/auth/pkg/user_v1"
)

// Delete удаляет клиента из базы
func (s *Server) Delete(ctx context.Context, req *userdesc.DeleteRequest) (*emptypb.Empty, error) {
	ctx = logger.AssignLogger(ctx, s.logger)

	log := s.logger.With(slog.String("method", "grpc-server.Delete"))
	log.Debug("called", slog.Int64("id", req.GetId()))

	err := s.srv.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
