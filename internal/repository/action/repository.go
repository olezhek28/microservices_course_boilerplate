package action

import (
	"context"

	"github.com/neracastle/auth/internal/repository/action/postgres/model"
)

// Repository хранилище действий клиента
type Repository interface {
	Save(context.Context, model.ActionDTO) error
}
