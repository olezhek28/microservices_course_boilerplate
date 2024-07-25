package action

import (
	"context"

	"github.com/neracastle/auth/internal/repository/action/postgres/model"
)

type Repository interface {
	Save(context.Context, model.ActionDTO) error
}
