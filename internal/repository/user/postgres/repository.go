package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/neracastle/go-libs/pkg/db"
	"github.com/neracastle/go-libs/pkg/sys/logger"
	"golang.org/x/exp/slog"

	domain "github.com/neracastle/auth/internal/domain/user"
	"github.com/neracastle/auth/internal/repository/user"
	pgmodel "github.com/neracastle/auth/internal/repository/user/postgres/model"
)

var _ user.Repository = (*repo)(nil)

type repo struct {
	conn db.Client
}

// New новый экземпляр репозитория pg
func New(conn db.Client) user.Repository {
	instance := &repo{conn: conn}

	return instance
}

func (r *repo) Save(ctx context.Context, user *domain.User) error {
	log := logger.GetLogger(ctx)
	log = log.With(slog.String("method", "repository.user.postgres.Save"))
	dto := FromDomainToRepo(user)

	q := db.Query{Name: "Save", QueryRaw: "INSERT INTO auth.users(email, password, name, role) VALUES ($1, $2, $3, $4) RETURNING id"}

	err := r.conn.DB().QueryRow(ctx, q,
		dto.Email,
		dto.Password,
		dto.Name,
		dto.IsAdmin).Scan(&user.ID)
	if err != nil {
		log.Error("failed to save user in db", slog.String("error", err.Error()))
		return err
	}

	log.Debug("saved user in db", slog.Int64("id", user.ID))

	return nil
}

func (r *repo) Update(ctx context.Context, user *domain.User) error {
	log := logger.GetLogger(ctx)
	dto := FromDomainToRepo(user)

	q := db.Query{Name: "Update", QueryRaw: "UPDATE auth.users SET name = $1, email = $2, password = $3, role = $4, updated_at = now() WHERE id = $5"}
	_, err := r.conn.DB().Exec(ctx, q,
		dto.Name,
		dto.Email,
		dto.Password,
		dto.IsAdmin,
		dto.ID)
	if err != nil {
		log.Error("failed to update user in db", slog.String("error", err.Error()), slog.String("method", "repository.user.postgres.Update"))
		return err
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	log := logger.GetLogger(ctx)
	q := db.Query{Name: "Delete", QueryRaw: "DELETE FROM auth.users WHERE id = $1"}
	_, err := r.conn.DB().Exec(ctx, q, id)
	if err != nil {
		log.Error("failed to delete user", slog.String("error", err.Error()), slog.String("method", "repository.user.postgres.Delete"))
		return err
	}

	return nil
}

func (r *repo) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	log := logger.GetLogger(ctx)
	log = log.With(slog.String("method", "repository.user.postgres.GetById"), slog.Int64("user_id", id))

	q := db.Query{Name: "GetById", QueryRaw: `SELECT id, email, password, name, role, created_at 
									  		FROM auth.users 
									  		WHERE id = $1`}
	res, err := r.conn.DB().Query(ctx, q, id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrUserNotFound
		}

		log.Error("failed to get user from db", slog.String("error", err.Error()))

		return nil, err
	}

	dto, err := pgx.CollectOneRow(res, pgx.RowToStructByName[pgmodel.UserDTO])
	if err != nil {
		return nil, err
	}

	userAggr := FromRepoToDomain(dto)

	return userAggr, nil
}

func (r *repo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	log := logger.GetLogger(ctx)
	log = log.With(slog.String("method", "repository.user.postgres.GetByEmail"), slog.String("email", email))

	q := db.Query{Name: "GetById", QueryRaw: `SELECT id, email, password, name, role, created_at 
									  		FROM auth.users 
									  		WHERE email = $1`}
	res, err := r.conn.DB().Query(ctx, q, email)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrUserNotFound
		}

		log.Error("failed to get user from db", slog.String("error", err.Error()))

		return nil, err
	}

	dto, err := pgx.CollectOneRow(res, pgx.RowToStructByName[pgmodel.UserDTO])
	if err != nil {
		return nil, err
	}

	userAggr := FromRepoToDomain(dto)

	return userAggr, nil
}
