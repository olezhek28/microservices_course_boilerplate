package user

import (
	"time"
)

// User доменный агрегат пользователя в системе
type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
	IsAdmin  bool
	RegDate  time.Time
}

// ChangeEmail меняет почту юзера
func (u *User) ChangeEmail(email string) error {
	if email == "" {
		return ErrEmptyEmail
	}

	u.Email = email
	return nil
}

// ChangePassword меняет пароль
func (u *User) ChangePassword(password string) error {
	if password == "" {
		return ErrEmptyPwd
	}

	u.Password = password
	return nil
}

// NewUser создает нового пользователя
func NewUser(email string, password string, name string) (*User, error) {
	if email == "" {
		return nil, ErrEmptyEmail
	}

	if password == "" {
		return nil, ErrEmptyPwd
	}

	return &User{
		Name:     name,
		Password: password,
		Email:    email,
		RegDate:  time.Now(),
	}, nil
}

// NewAdmin Создает нового пользователя с ролью Admin
func NewAdmin(email string, password string, name string) (*User, error) {
	usr, err := NewUser(email, password, name)
	if err != nil {
		return usr, err
	}

	usr.IsAdmin = true
	return usr, nil
}
