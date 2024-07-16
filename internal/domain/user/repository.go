package user

import "errors"

type Repository interface {
	Save(userModel User) error
	Update(userModel User) error
	Delete(id int64) error
	GetById(id int64) (User, error)
	FindByUsername(username string) (User, error)
}

var (
	RepoErrUserNotFound = errors.New("пользователь не найден")
)
