package models

import (
	"github.com/neracastle/auth/internal/domain/user"
)

func FromDomainToUsecase(dbUser *user.User) UserDTO {
	return UserDTO{
		Id:        dbUser.Id,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
		Name:      dbUser.Name,
		IsAdmin:   dbUser.IsAdmin,
		CreatedAt: dbUser.RegDate,
	}
}
