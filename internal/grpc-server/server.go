package grpc_server

import (
	"context"
	"log"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	userdesc "github.com/neracastle/auth/pkg/user_v1"
)

// Server GRPC сервер с ручками сервиса auth
type Server struct {
	userdesc.UnimplementedUserV1Server
}

func (s *Server) Create(ctx context.Context, req *userdesc.CreateRequest) (*userdesc.CreateResponse, error) {
	log.Printf("called Create method with req: %v", req)

	return &userdesc.CreateResponse{Id: gofakeit.Int64()}, nil
}

func (s *Server) Get(ctx context.Context, req *userdesc.GetRequest) (*userdesc.GetResponse, error) {
	log.Printf("called Get method with req: %v", req)

	return &userdesc.GetResponse{
		Id:        req.GetId(),
		Name:      gofakeit.BeerName(),
		Email:     gofakeit.Email(),
		Role:      userdesc.Role_USER,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

func (s *Server) Update(ctx context.Context, req *userdesc.UpdateRequest) (*emptypb.Empty, error) {
	if req.GetRole() == userdesc.Role_UNKNOWN {
		log.Printf("called Update method with empty role: %v", req)
	} else {
		log.Printf("called Update method with setted role: %v", req)
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) Delete(ctx context.Context, req *userdesc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("called Delete method with req: %v", req)

	return &emptypb.Empty{}, nil
}
