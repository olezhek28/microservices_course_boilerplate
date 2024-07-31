package postgres

import (
	"context"

	"github.com/neracastle/go-libs/pkg/db"
	"github.com/neracastle/go-libs/pkg/sys/logger"
	"golang.org/x/exp/slog"

	"github.com/neracastle/auth/internal/repository/action"
	"github.com/neracastle/auth/internal/repository/action/postgres/model"
)

var _ action.Repository = (*repo)(nil)

type repo struct {
	conn db.Client
}

// New новый экземпляр репозитария
func New(conn db.Client) action.Repository {
	instance := &repo{conn: conn}

	return instance
}

func (r *repo) Save(ctx context.Context, dto model.ActionDTO) error {
	log := logger.GetLogger(ctx)
	log = log.With(slog.String("method", "repository.postgres.Save"))

	q := db.Query{Name: "Save", QueryRaw: "INSERT INTO auth.user_actions(user_id, name, old_value, new_value) VALUES ($1, $2, $3, $4)"}

	_, err := r.conn.DB().Exec(ctx, q,
		dto.UserID,
		dto.Name,
		dto.OldValue,
		dto.NewValue)
	if err != nil {
		log.Error("failed to save user action in db", slog.String("error", err.Error()))
		return err
	}

	return nil
}
