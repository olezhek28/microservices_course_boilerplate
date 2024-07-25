package usecases

import (
	db "github.com/neracastle/auth/internal/client"
	"github.com/neracastle/auth/internal/repository/action"
	"github.com/neracastle/auth/internal/repository/user"
)

type Service struct {
	urepo user.Repository
	arepo action.Repository
	db    db.DB
}

func NewService(urepo user.Repository, arepo action.Repository, db db.DB) *Service {
	return &Service{
		urepo: urepo,
		arepo: arepo,
		db:    db,
	}
}
