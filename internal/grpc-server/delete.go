package grpc_server

import (
	"context"

	userdesc "github.com/neracastle/auth/pkg/user_v1"
)

// Delete удаляет клиента из базы
func (s *Server) Delete(ctx context.Context, req *userdesc.DeleteRequest) (*userdesc.DeleteResponse, error) {
	err := s.srv.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &userdesc.DeleteResponse{}, nil
}
