package usecases

import (
	"context"

	"golang.org/x/exp/slog"

	"github.com/neracastle/auth/internal/app/logger"
)

func (s *Service) Delete(ctx context.Context, userId int64) error {
	log := logger.GetLogger(ctx)
	log.Debug("called", slog.String("method", "usecases.Delete"), slog.Int64("user_id", userId))

	return s.urepo.Delete(ctx, userId)
}
