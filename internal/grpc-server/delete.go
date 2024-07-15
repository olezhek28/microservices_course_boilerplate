package grpc_server

import (
	"context"

	"golang.org/x/exp/slog"
	"google.golang.org/protobuf/types/known/emptypb"

	userdesc "github.com/neracastle/auth/pkg/user_v1"
)

// Delete удаляет клиента из базы
func (s *Server) Delete(ctx context.Context, req *userdesc.DeleteRequest) (*emptypb.Empty, error) {
	log := s.GetLogger()
	log = log.With(slog.String("method", "grpc-server.Delete"))

	log.Debug("called", slog.Int64("id", req.GetId()))

	//далее скроется за репо-слоем
	res, err := s.pgcon.Exec(ctx, "DELETE FROM auth.users WHERE id=$1", req.GetId())
	if err != nil {
		log.Error("failed to delete user", slog.String("error", err.Error()))
		return nil, err
	}

	log.Debug("deleted user",
		slog.Int64("id", req.GetId()),
		slog.Int64("rows", res.RowsAffected()))

	return &emptypb.Empty{}, nil
}
