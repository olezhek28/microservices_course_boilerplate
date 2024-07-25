package models

type CreateDTO struct {
	Email           string
	Password        string
	PasswordConfirm string
	Name            string
	IsAdmin         bool
}
