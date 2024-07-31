package models

// CreateDTO входные данные для создания пользователя
type CreateDTO struct {
	Email           string
	Password        string
	PasswordConfirm string
	Name            string
	IsAdmin         bool
}
