package grpc_server

import (
	"context"
	"errors"
	"time"

	"golang.org/x/exp/slog"

	"github.com/neracastle/auth/internal/domain/user"
	userdesc "github.com/neracastle/auth/pkg/user_v1"
)

// Create регистрирует нового пользователя
func (s *Server) Create(ctx context.Context, req *userdesc.CreateRequest) (*userdesc.CreateResponse, error) {
	log := s.GetLogger()
	log = log.With(slog.String("method", "grpc-server.Create"))

	log.Debug("called", slog.Any("req", req))

	//здесь в дальнейшем будет app-сервис и репозиторий
	if req.Password != req.PasswordConfirm {
		log.Debug("passwords do not match",
			slog.String("pwd", req.Password),
			slog.String("pwd_confirm", req.PasswordConfirm))

		return nil, errors.New("пароли не совпадают")
	}

	if req.GetRole() == userdesc.Role_UNKNOWN {
		log.Error("role is unknown")
		return nil, errors.New("роль не задана")
	}

	var err error
	var newUser user.User
	if req.GetRole() == userdesc.Role_USER {
		newUser, err = user.NewUser(req.GetName(), req.GetPassword(), req.GetEmail())
	} else {
		newUser, err = user.NewAdmin(req.GetName(), req.GetPassword(), req.GetEmail())
	}

	if err != nil {
		log.Error("failed to create user", slog.String("error", err.Error()))
		return nil, err
	}

	//далее все скроется за репо слоем
	var id int64
	var created time.Time
	var role = userdesc.Role_USER

	if newUser.IsAdmin() {
		role = userdesc.Role_ADMIN
	}

	err = s.pgcon.QueryRow(ctx, "INSERT INTO auth.users(email, password, name, role) VALUES ($1, $2, $3, $4) RETURNING id, created_at",
		newUser.GetEmail(),
		newUser.GetPassword(),
		newUser.GetName(),
		role).Scan(&id, &created)
	if err != nil {
		log.Error("failed to save user in db", slog.String("error", err.Error()))
		return nil, err
	}

	log.Debug("saved user in db",
		slog.Int64("id", id),
		slog.String("time", created.String()))

	return &userdesc.CreateResponse{Id: id}, nil
}
