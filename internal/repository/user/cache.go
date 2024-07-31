package user

import (
	"context"
	"errors"

	domain "github.com/neracastle/auth/internal/domain/user"
)

// Cache кэш пользователей
type Cache interface {
	Save(context.Context, *domain.User) error
	GetByID(ctx context.Context, id int64) (*domain.User, error)
}

var (
	// ErrUserNotCached если пользователя нет в кэше
	ErrUserNotCached = errors.New("пользователь не найден")
)
