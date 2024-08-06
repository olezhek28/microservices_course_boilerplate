package grpc_server

import (
	"context"
	"errors"

	userdesc "github.com/neracastle/auth/pkg/user_v1"
)

// Create регистрирует нового пользователя
func (s *Server) Create(ctx context.Context, req *userdesc.CreateRequest) (*userdesc.CreateResponse, error) {
	if req.GetRole() == userdesc.Role_UNKNOWN {
		return nil, errors.New("роль не задана")
	}

	apiDto := FromGrpcToCreateUsecase(req)
	userID, err := s.srv.Create(ctx, apiDto)

	if err != nil {
		return nil, err
	}

	return &userdesc.CreateResponse{Id: userID}, nil
}
