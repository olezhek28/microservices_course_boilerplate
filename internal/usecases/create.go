package usecases

import (
	"context"
	"errors"

	"golang.org/x/exp/slog"

	"github.com/neracastle/auth/internal/app/logger"
	domain "github.com/neracastle/auth/internal/domain/user"
	def "github.com/neracastle/auth/internal/usecases/models"
)

func (s *Service) Create(ctx context.Context, req def.CreateDTO) (int64, error) {
	log := logger.GetLogger(ctx)
	log.Debug("called", slog.String("method", "usecases.Create"))

	var err error
	var newUser *domain.User

	if req.Password != req.PasswordConfirm {
		return 0, errors.New("пароли не совпадают")
	}

	if req.IsAdmin {
		newUser, err = domain.NewAdmin(req.Email, req.Password, req.Name)
	} else {
		newUser, err = domain.NewUser(req.Email, req.Password, req.Name)
	}

	if err != nil {
		return 0, err
	}

	err = s.urepo.Save(ctx, newUser)
	if err != nil {
		log.Error("failed to create user", slog.String("error", err.Error()), slog.String("method", "usecases.Create"))
		return 0, err
	}

	return newUser.Id, nil
}
