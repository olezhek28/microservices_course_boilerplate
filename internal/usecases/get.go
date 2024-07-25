package usecases

import (
	"context"

	"golang.org/x/exp/slog"

	"github.com/neracastle/auth/internal/app/logger"
	"github.com/neracastle/auth/internal/usecases/models"
)

func (s *Service) Get(ctx context.Context, userId int64) (models.UserDTO, error) {
	log := logger.GetLogger(ctx)
	log.Debug("called", slog.String("method", "usecases.Get"), slog.Int64("user_id", userId))

	dbUser, err := s.urepo.GetById(ctx, userId)

	return models.FromDomainToUsecase(dbUser), err
}
