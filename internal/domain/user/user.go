package user

import (
	"errors"
	"time"
)

var _ User = (*user)(nil)

// User агрегат доменного слоя, представляет пользователя и его бизнес-логику
type User interface {
	GetID() int64
	GetName() string
	GetEmail() string
	GetPassword() string
	IsAdmin() bool
	ChangeName(newName string)
	ChangeEmail(newEmail string) error
	ChangePassword(newPassword string) error
	SetIsAdmin()
	SetIsUser()
}

type user struct {
	id        int64
	name      string
	email     string
	password  string
	isAdmin   bool
	createdAt time.Time
}

func (u *user) GetID() int64 {
	return u.id
}

func (u *user) GetName() string {
	return u.name
}

func (u *user) GetEmail() string {
	return u.email
}

func (u *user) GetPassword() string {
	return u.password
}

func (u *user) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *user) IsAdmin() bool {
	return u.isAdmin
}

func (u *user) ChangeName(name string) {
	u.name = name
}

func (u *user) ChangeEmail(email string) error {
	if email == "" {
		return errors.New("email не может быть пустым")
	}

	u.email = email
	return nil
}

func (u *user) ChangePassword(password string) error {
	if password == "" {
		return errors.New("пароль не может быть пустым")
	}

	u.password = password
	return nil
}

func (u *user) SetIsAdmin() {
	u.isAdmin = true
}

func (u *user) SetIsUser() {
	u.isAdmin = false
}

func NewUser(name string, password string, email string) (User, error) {
	if email == "" {
		return &user{}, errors.New("email не может быть пустым")
	}

	if password == "" {
		return &user{}, errors.New("пароль не может быть пустым")
	}

	return &user{
		name:      name,
		password:  password,
		email:     email,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(name string, password string, email string) (User, error) {
	usr, err := NewUser(name, password, email)
	if err != nil {
		return usr, err
	}

	usr.SetIsAdmin()
	return usr, nil
}
