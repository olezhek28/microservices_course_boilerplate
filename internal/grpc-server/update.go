package grpc_server

import (
	"context"

	userdesc "github.com/neracastle/auth/pkg/user_v1"
)

// Update обновляет данные клиента
func (s *Server) Update(ctx context.Context, req *userdesc.UpdateRequest) (*userdesc.UpdateResponse, error) {
	dto := FromGrpcToUpdateUsecase(req)
	err := s.srv.Update(ctx, dto)
	if err != nil {
		return nil, err
	}

	return &userdesc.UpdateResponse{}, nil
}
