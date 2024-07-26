package app

import (
	"context"
	"log"

	db "github.com/neracastle/auth/internal/client"
	"github.com/neracastle/auth/internal/client/pg"
	"github.com/neracastle/auth/internal/config"
	"github.com/neracastle/auth/internal/repository/action"
	acRepo "github.com/neracastle/auth/internal/repository/action/postgres"
	"github.com/neracastle/auth/internal/repository/user"
	uRepo "github.com/neracastle/auth/internal/repository/user/postgres"
	"github.com/neracastle/auth/internal/usecases"
)

type serviceProvider struct {
	conf           *config.Config
	usecaseService *usecases.Service
	usersRepo      user.Repository
	actionsRepo    action.Repository
	dbc            db.Client
	txm            db.TxManager
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) Config() config.Config {
	if sp.conf == nil {
		cfg := config.MustLoad()
		sp.conf = &cfg
	}

	return *sp.conf
}

func (sp *serviceProvider) DbClient(ctx context.Context) db.Client {
	if sp.dbc == nil {
		client, err := pg.NewClient(ctx, sp.Config().Postgres.DSN())
		if err != nil {
			log.Fatalf("failed to connect to pg: %v", err)
		}

		err = client.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("failed ping to pg: %v", err)
		}

		sp.dbc = client
	}

	return sp.dbc
}

func (sp *serviceProvider) UsersRepository(ctx context.Context) user.Repository {
	if sp.usersRepo == nil {
		sp.usersRepo = uRepo.New(sp.DbClient(ctx))
	}

	return sp.usersRepo
}

func (sp *serviceProvider) ActionsRepository(ctx context.Context) action.Repository {
	if sp.actionsRepo == nil {
		sp.actionsRepo = acRepo.New(sp.DbClient(ctx))
	}

	return sp.actionsRepo
}

func (sp *serviceProvider) UsersService(ctx context.Context) *usecases.Service {
	if sp.usecaseService == nil {
		sp.usecaseService = usecases.NewService(
			sp.UsersRepository(ctx),
			sp.ActionsRepository(ctx),
			sp.DbClient(ctx).DB())
	}

	return sp.usecaseService
}
