package usecases

import (
	"context"
	"time"

	"golang.org/x/exp/slog"

	"github.com/neracastle/auth/internal/app/logger"
	"github.com/neracastle/auth/internal/repository/action/postgres/model"
	def "github.com/neracastle/auth/internal/usecases/models"
)

func (s *Service) Update(ctx context.Context, user def.UpdateDTO) error {
	log := logger.GetLogger(ctx)
	log.Debug("called", slog.String("method", "usecases.Update"), slog.Int64("user_id", user.Id))

	dbUser, err := s.urepo.GetById(ctx, user.Id)

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
		err = s.urepo.Update(ctx, dbUser)
		if err != nil {
			return err
		}

		err = s.arepo.Save(ctx, model.ActionDTO{
			UserId:    dbUser.Id,
			Name:      "ChangeEmail",
			OldValue:  oldEmail,
			NewValue:  dbUser.Email,
			CreatedAt: time.Now(),
		})

		return err
	})

	return err
}
