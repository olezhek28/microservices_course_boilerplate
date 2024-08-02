package usecases

import (
	"context"
	"time"

	"github.com/neracastle/go-libs/pkg/sys/logger"
	"golang.org/x/exp/slog"

	"github.com/neracastle/auth/internal/repository/action/postgres/model"
	userRepo "github.com/neracastle/auth/internal/repository/user"
	def "github.com/neracastle/auth/internal/usecases/models"
)

// Update обновляет данные пользователя
func (s *Service) Update(ctx context.Context, user def.UpdateDTO) error {
	log := logger.GetLogger(ctx)
	log.Debug("called", slog.String("method", "usecases.Update"), slog.Int64("user_id", user.ID))

	dbUser, err := s.usersRepo.Get(ctx, userRepo.SearchFilter{ID: user.ID})

	if err != nil {
		return err
	}

	dbUser.Name = user.Name

	switch user.Role {
	case 1:
		dbUser.IsAdmin = false
	case 2:
		dbUser.IsAdmin = true
	}

	oldEmail := dbUser.Email
	err = dbUser.ChangeEmail(user.Email)
	if err != nil {
		return err
	}

	err = s.db.ReadCommitted(ctx, func(ctx context.Context) error {
		err = s.usersRepo.Update(ctx, dbUser)
		if err != nil {
			return err
		}

		err = s.actionsRepo.Save(ctx, model.ActionDTO{
			UserID:    dbUser.ID,
			Name:      "ChangeEmail",
			OldValue:  oldEmail,
			NewValue:  dbUser.Email,
			CreatedAt: time.Now(),
		})

		return err
	})

	return err
}
