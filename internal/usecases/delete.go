package usecases

import (
	"context"

	"github.com/neracastle/go-libs/pkg/sys/logger"
	"golang.org/x/exp/slog"
)

// Delete удаляет пользователя
func (s *Service) Delete(ctx context.Context, userID int64) error {
	log := logger.GetLogger(ctx)
	log.Debug("called", slog.String("method", "usecases.Delete"), slog.Int64("user_id", userID))

	return s.usersRepo.Delete(ctx, userID)
}
