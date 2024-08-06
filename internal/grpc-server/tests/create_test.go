package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/neracastle/go-libs/pkg/sys/logger"
	"github.com/stretchr/testify/require"

	grpc_server "github.com/neracastle/auth/internal/grpc-server"
	"github.com/neracastle/auth/internal/usecases"
	"github.com/neracastle/auth/internal/usecases/mocks"
	"github.com/neracastle/auth/internal/usecases/models"
	"github.com/neracastle/auth/pkg/user_v1"
)

func TestCreate(t *testing.T) {
	type caseArgs struct {
		ctx context.Context
		req *user_v1.CreateRequest
	}

	var (
		mc  = minimock.NewController(t)
		ctx = context.Background()

		pwd       = gofakeit.Password(true, true, true, false, false, 8)
		createDTO = models.CreateDTO{
			Email:           gofakeit.Email(),
			Password:        pwd,
			PasswordConfirm: pwd,
			Name:            gofakeit.Name(),
			IsAdmin:         false,
		}

		wrongRoleErr = errors.New("роль не задана")
	)

	tests := []struct {
		name            string
		args            caseArgs
		want            *user_v1.CreateResponse
		err             error
		userServiceMock func(mc *minimock.Controller) usecases.UserService
	}{
		{
			name: "success",
			args: caseArgs{
				ctx: ctx,
				req: &user_v1.CreateRequest{
					Name:            createDTO.Name,
					Email:           createDTO.Email,
					Password:        pwd,
					PasswordConfirm: pwd,
					Role:            user_v1.Role_USER,
				},
			},
			want: &user_v1.CreateResponse{
				Id: 1,
			},
			err: nil,
			userServiceMock: func(mc *minimock.Controller) usecases.UserService {
				mockedSrv := mocks.NewUserServiceMock(mc)
				lctx := logger.AssignLogger(ctx, nil)
				mockedSrv.CreateMock.Expect(lctx, createDTO).Return(1, nil)

				return mockedSrv
			},
		},
		{
			name: "wrong role",
			args: caseArgs{
				ctx: ctx,
				req: &user_v1.CreateRequest{
					Name:            createDTO.Name,
					Email:           createDTO.Email,
					Password:        pwd,
					PasswordConfirm: pwd,
					Role:            user_v1.Role_UNKNOWN,
				},
			},
			want: nil,
			err:  wrongRoleErr,
			userServiceMock: func(mc *minimock.Controller) usecases.UserService {
				return mocks.NewUserServiceMock(mc)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockUserService := tt.userServiceMock(mc)
			grpcService := grpc_server.NewServer(mockUserService)
			res, err := grpcService.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.want, res)
			require.Equal(t, tt.err, err)
		})
	}
}
