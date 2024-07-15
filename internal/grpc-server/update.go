package grpc_server

import (
	"context"
	"errors"

	"golang.org/x/exp/slog"
	"google.golang.org/protobuf/types/known/emptypb"

	userdesc "github.com/neracastle/auth/pkg/user_v1"
)

// Update обновляет данные клиента
func (s *Server) Update(ctx context.Context, req *userdesc.UpdateRequest) (*emptypb.Empty, error) {
	log := s.GetLogger()
	log = log.With(slog.String("method", "grpc-server.Update"))

	log.Debug("called", slog.Any("req", req))

	//здесь в дальнейшем будет app-сервис и репозиторий
	if req.GetRole() == userdesc.Role_UNKNOWN {
		log.Error("role is unknown")
		return nil, errors.New("роль не задана")
	}

	if req.Email != nil && req.GetEmail() == "" {
		log.Error("email is empty")
		return nil, errors.New("email не может быть пустым")
	}

	//доменные модельки пока не сохраняем, они будут идти в методы репозитория
	res, err := s.pgcon.Exec(ctx, "UPDATE auth.users SET name = $1, email = $2, role = $3, updated_at = now() WHERE id = $4",
		req.GetName(),
		req.GetEmail(),
		req.GetRole(),
		req.GetId())
	if err != nil {
		log.Error("failed to update user in db", slog.String("error", err.Error()))
		return nil, err
	}

	log.Debug("updated user in db",
		slog.Int64("id", req.Id),
		slog.Int64("affected rows", res.RowsAffected()))

	return &emptypb.Empty{}, nil
}
