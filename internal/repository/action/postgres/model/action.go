package model

import (
	"time"
)

type ActionDTO struct {
	UserId    int64     `db:"user_id"`
	Name      string    `db:"name"`
	OldValue  string    `db:"old_value"`
	NewValue  string    `db:"new_value"`
	CreatedAt time.Time `db:"created_at"`
}
