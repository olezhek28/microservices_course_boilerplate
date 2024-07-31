package models

import "time"

// UserDTO модель данных пользователя на сервисном слое
type UserDTO struct {
	ID        int64
	Email     string
	Password  string
	Name      string
	IsAdmin   bool
	CreatedAt time.Time
}
