package grpc_server

import (
	"context"

	userdesc "github.com/neracastle/auth/pkg/user_v1"
)

// Get возвращает данные клиента
func (s *Server) Get(ctx context.Context, req *userdesc.GetRequest) (*userdesc.GetResponse, error) {
	dbUser, err := s.srv.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	rsp := FromUsecaseToGetResponse(dbUser)

	return rsp, nil
}
