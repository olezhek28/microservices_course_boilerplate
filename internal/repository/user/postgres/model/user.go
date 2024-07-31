package model

import (
	"database/sql"
	"time"
)

// UserDTO модель для представления в pg
type UserDTO struct {
	ID        int64          `db:"id"`
	Email     string         `db:"email"`
	Password  string         `db:"password"`
	Name      sql.NullString `db:"name"`
	IsAdmin   int8           `db:"role"`
	CreatedAt time.Time      `db:"created_at"`
}
