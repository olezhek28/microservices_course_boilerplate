package model

// UserDTO модель юзера для хранения в редисе
type UserDTO struct {
	ID        int64  `redis:"id"`
	Email     string `redis:"email"`
	Name      string `redis:"name"`
	IsAdmin   int8   `redis:"is_admin"`
	CreatedAt int64  `redis:"created_at"`
}
