package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/neracastle/go-libs/pkg/redis"

	domain "github.com/neracastle/auth/internal/domain/user"
	"github.com/neracastle/auth/internal/repository/user"
	"github.com/neracastle/auth/internal/repository/user/redis/model"
)

var _ user.Cache = (*repo)(nil)

type repo struct {
	client redis.Client
}

// New новый экземпляр клиента
func New(client redis.Client) user.Cache {
	return &repo{
		client: client,
	}
}

func (r *repo) Save(ctx context.Context, d *domain.User, ttl time.Duration) error {
	dto := FromDomainToRepo(d)
	err := r.client.HSetMap(ctx, r.getKey(dto.ID), dto)
	if err != nil {
		return err
	}

	err = r.client.Expire(ctx, r.getKey(dto.ID), ttl)
	if err != nil {
		_ = r.client.Del(ctx, r.getKey(dto.ID))
		return err
	}

	return nil
}

// GetByID возвращает пользователя по ID. Если не найден, вернется user.ErrUserNotCached
func (r *repo) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	exist, err := r.client.Exist(ctx, r.getKey(id))
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, user.ErrUserNotCached
	}

	var dto model.UserDTO
	err = r.client.HGetAll(ctx, r.getKey(id), &dto)
	if err != nil {
		return nil, err
	}

	return FromRepoToDomain(dto), nil
}

func (r *repo) getKey(id int64) string {
	return fmt.Sprintf("user:%d", id)
}
