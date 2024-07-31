package grpc_server

import (
	"context"

	"github.com/neracastle/go-libs/pkg/sys/logger"
	"golang.org/x/exp/slog"
	"google.golang.org/protobuf/types/known/emptypb"

	userdesc "github.com/neracastle/auth/pkg/user_v1"
)

// Update обновляет данные клиента
func (s *Server) Update(ctx context.Context, req *userdesc.UpdateRequest) (*emptypb.Empty, error) {
	ctx = logger.AssignLogger(ctx, s.logger)

	log := s.logger.With(slog.String("method", "grpc-server.Update"))
	log.Debug("called", slog.Any("req", req))

	dto := FromGrpcToUpdateUsecase(req)
	err := s.srv.Update(ctx, dto)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
