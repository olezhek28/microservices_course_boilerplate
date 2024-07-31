package model

import (
	"time"
)

// ActionDTO модель события действий пользователя
type ActionDTO struct {
	UserID    int64     `db:"user_id"`
	Name      string    `db:"name"`
	OldValue  string    `db:"old_value"`
	NewValue  string    `db:"new_value"`
	CreatedAt time.Time `db:"created_at"`
}
