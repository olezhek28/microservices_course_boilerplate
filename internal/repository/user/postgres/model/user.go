package model

import (
	"database/sql"
	"time"
)

type UserDTO struct {
	Id        int64          `db:"id"`
	Email     string         `db:"email"`
	Password  string         `db:"password"`
	Name      sql.NullString `db:"name"`
	IsAdmin   int8           `db:"role"`
	CreatedAt time.Time      `db:"created_at"`
}
