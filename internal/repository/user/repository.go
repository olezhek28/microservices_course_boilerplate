package user

import (
	"context"
	"errors"

	domain "github.com/neracastle/auth/internal/domain/user"
)

// Repository репозитарий пользователей
type Repository interface {
	Save(context.Context, *domain.User) error
	Update(context.Context, *domain.User) error
	Delete(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
}

var (
	// ErrUserNotFound пользователь отсутствует в хранилище
	ErrUserNotFound = errors.New("пользователь не найден")
)
