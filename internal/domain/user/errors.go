package user

import "errors"

// ErrEmptyEmail если задан пустой email
var ErrEmptyEmail = errors.New("email не может быть пустым")

// ErrEmptyPwd если задан пустой пароль
var ErrEmptyPwd = errors.New("пароль не может быть пустым")
