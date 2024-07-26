package models

import "time"

type UserDTO struct {
	Id        int64
	Email     string
	Password  string
	Name      string
	IsAdmin   bool
	CreatedAt time.Time
}
