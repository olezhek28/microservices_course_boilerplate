package postgres

import (
	domain "github.com/neracastle/auth/internal/domain/user"
	pg_repo "github.com/neracastle/auth/internal/repository/user/postgres/model"
)

// FromDomainToRepo преобразует доменную сущность в дто хранилища
func FromDomainToRepo(user *domain.User) pg_repo.UserDTO {
	dto := pg_repo.UserDTO{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
		IsAdmin:  0,
	}

	if user.IsAdmin {
		dto.IsAdmin = 1
	}

	if user.Name != "" {
		_ = dto.Name.Scan(user.Name)
	}

	return dto
}

// FromRepoToDomain преобразует дто хранилища в доменную сущность
func FromRepoToDomain(dto pg_repo.UserDTO) *domain.User {
	return &domain.User{
		ID:       dto.ID,
		Email:    dto.Email,
		Password: dto.Password,
		Name:     dto.Name.String,
		RegDate:  dto.CreatedAt,
		IsAdmin:  dto.IsAdmin > 0,
	}
}
