package grpc_server

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"golang.org/x/exp/slog"
	"google.golang.org/protobuf/types/known/timestamppb"

	userdesc "github.com/neracastle/auth/pkg/user_v1"
)

// Get возвращает данные клиента
func (s *Server) Get(ctx context.Context, req *userdesc.GetRequest) (*userdesc.GetResponse, error) {
	log := s.GetLogger()
	log = log.With(slog.String("method", "grpc-server.Get"))

	log.Debug("called", slog.Any("id", req.GetId()))

	//далее скроется за репо-слоем
	res := s.pgcon.QueryRow(ctx, `SELECT id, email, name, role, created_at, updated_at 
									  FROM auth.users 
									  WHERE id = $1`, req.GetId())

	var created time.Time
	var updated sql.NullTime
	var rsp userdesc.GetResponse
	err := res.Scan(&rsp.Id, &rsp.Email, &rsp.Name, &rsp.Role, &created, &updated)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Debug("user not found in db", slog.Int64("id", req.GetId()))
		} else {
			log.Error("failed to get user in db", slog.String("error", err.Error()))
		}

		return nil, err
	}

	rsp.CreatedAt = timestamppb.New(created)

	if updated.Valid {
		rsp.UpdatedAt = timestamppb.New(updated.Time)
	}

	return &rsp, nil
}
