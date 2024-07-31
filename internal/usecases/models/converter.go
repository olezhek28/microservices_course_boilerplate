package models

import (
	"github.com/neracastle/auth/internal/domain/user"
)

// FromDomainToUsecase преобразует доменную сущность в дто из сервисного слоя
func FromDomainToUsecase(dbUser *user.User) UserDTO {
	return UserDTO{
		ID:        dbUser.ID,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
		Name:      dbUser.Name,
		IsAdmin:   dbUser.IsAdmin,
		CreatedAt: dbUser.RegDate,
	}
}
