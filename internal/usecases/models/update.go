package models

type UpdateDTO struct {
	Id       int64
	Email    string
	Password string
	Name     string
	Role     uint8
}
