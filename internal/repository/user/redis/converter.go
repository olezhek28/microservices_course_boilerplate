package redis

import (
	"time"

	domain "github.com/neracastle/auth/internal/domain/user"
	"github.com/neracastle/auth/internal/repository/user/redis/model"
)

// FromDomainToRepo преобразует доменную сущность в дто хранилища
func FromDomainToRepo(user *domain.User) model.UserDTO {
	dto := model.UserDTO{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		IsAdmin:   0,
		CreatedAt: user.RegDate.Unix(),
	}

	if user.IsAdmin {
		dto.IsAdmin = 1
	}

	return dto
}

// FromRepoToDomain преобразует дто хранилища в доменную сущность
func FromRepoToDomain(dto model.UserDTO) *domain.User {
	return &domain.User{
		ID:      dto.ID,
		Email:   dto.Email,
		Name:    dto.Name,
		RegDate: time.Unix(dto.CreatedAt, 0),
		IsAdmin: dto.IsAdmin > 0,
	}
}
