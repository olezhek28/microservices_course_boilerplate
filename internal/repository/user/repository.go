package user

import (
	"context"
	"errors"

	domain "github.com/neracastle/auth/internal/domain/user"
)

type SearchFilter struct {
	ID    int64
	Email string
}

// Repository репозитарий пользователей
type Repository interface {
	Save(context.Context, *domain.User) error
	Update(context.Context, *domain.User) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, filter SearchFilter) (*domain.User, error)
}

var (
	// ErrUserNotFound пользователь отсутствует в хранилище
	ErrUserNotFound = errors.New("пользователь не найден")
)
