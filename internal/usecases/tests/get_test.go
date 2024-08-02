package tests

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/neracastle/go-libs/pkg/sys/logger"
	"github.com/stretchr/testify/require"

	domain "github.com/neracastle/auth/internal/domain/user"
	"github.com/neracastle/auth/internal/repository/user"
	"github.com/neracastle/auth/internal/repository/user/mocks"
	usecases2 "github.com/neracastle/auth/internal/usecases"
	usecases "github.com/neracastle/auth/internal/usecases/models"
)

func TestCreate(t *testing.T) {
	type caseArgs struct {
		ctx context.Context
		req user.SearchFilter
	}

	var (
		mc  = minimock.NewController(t)
		lg  = logger.SetupLogger("disable")
		ctx = logger.AssignLogger(context.Background(), lg)

		userAggr = &domain.User{
			ID:       gofakeit.Int64(),
			Email:    gofakeit.Email(),
			Password: gofakeit.Password(true, true, true, false, false, 8),
			Name:     gofakeit.Name(),
			IsAdmin:  gofakeit.Bool(),
			RegDate:  time.Now(),
		}
		userDTO = usecases.FromDomainToUsecase(userAggr)

		repoErr = errors.New("repo error")
	)

	tests := []struct {
		name           string
		args           caseArgs
		want           usecases.UserDTO
		err            error
		usersRepoMock  func(mc *minimock.Controller) user.Repository
		usersCacheMock func(mc *minimock.Controller) user.Cache
	}{
		{
			name: "Success. Not cached",
			args: caseArgs{
				ctx: ctx,
				req: user.SearchFilter{ID: userAggr.ID},
			},
			want: userDTO,
			err:  nil,
			usersRepoMock: func(mc *minimock.Controller) user.Repository {
				repoMock := mocks.NewRepositoryMock(mc)
				repoMock.GetMock.Expect(ctx, user.SearchFilter{ID: userAggr.ID}).Return(userAggr, nil)

				return repoMock
			},
			usersCacheMock: func(mc *minimock.Controller) user.Cache {
				cacheMock := mocks.NewCacheMock(mc)
				cacheMock.GetByIDMock.Expect(ctx, userAggr.ID).Return(nil, user.ErrUserNotCached)
				cacheMock.SaveMock.Expect(ctx, userAggr, time.Minute).Return(nil)

				return cacheMock
			},
		},
		{
			name: "Success. Cached",
			args: caseArgs{
				ctx: ctx,
				req: user.SearchFilter{ID: userAggr.ID},
			},
			want: userDTO,
			err:  nil,
			usersRepoMock: func(mc *minimock.Controller) user.Repository {
				repoMock := mocks.NewRepositoryMock(mc)
				//вызов репы не ожидаем, должен отработать только кэш
				return repoMock
			},
			usersCacheMock: func(mc *minimock.Controller) user.Cache {
				cacheMock := mocks.NewCacheMock(mc)
				cacheMock.GetByIDMock.Expect(ctx, userAggr.ID).Return(userAggr, nil)

				return cacheMock
			},
		},
		{
			name: "Failed. Repository Error",
			args: caseArgs{
				ctx: ctx,
				req: user.SearchFilter{ID: userAggr.ID},
			},
			want: usecases.UserDTO{},
			err:  repoErr,
			usersRepoMock: func(mc *minimock.Controller) user.Repository {
				repoMock := mocks.NewRepositoryMock(mc)
				repoMock.GetMock.Expect(ctx, user.SearchFilter{ID: userAggr.ID}).Return(nil, repoErr)
				return repoMock
			},
			usersCacheMock: func(mc *minimock.Controller) user.Cache {
				cacheMock := mocks.NewCacheMock(mc)
				cacheMock.GetByIDMock.Expect(ctx, userAggr.ID).Return(nil, user.ErrUserNotCached)

				return cacheMock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.usersRepoMock(mc)
			cache := tt.usersCacheMock(mc)

			srv := usecases2.NewService(repo, cache, nil, nil, time.Minute)
			res, err := srv.Get(tt.args.ctx, tt.args.req.ID)
			require.Equal(t, tt.want, res)
			require.Equal(t, tt.err, err)
		})
	}
}
