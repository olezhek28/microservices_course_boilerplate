package models

// UpdateDTO входные данные для запроса обновления юзера
type UpdateDTO struct {
	ID       int64
	Email    string
	Password string
	Name     string
	Role     uint8
}
