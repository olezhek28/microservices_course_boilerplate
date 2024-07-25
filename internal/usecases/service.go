package usecases

import (
	db "github.com/neracastle/auth/internal/client"
	"github.com/neracastle/auth/internal/repository/action"
	"github.com/neracastle/auth/internal/repository/user"
)

type Service struct {
	urepo user.Repository
	arepo action.Repository
	tx    db.TxManager
}

func NewService(urepo user.Repository, arepo action.Repository, tx db.TxManager) *Service {
	return &Service{
		urepo: urepo,
		arepo: arepo,
		tx:    tx,
	}
}
