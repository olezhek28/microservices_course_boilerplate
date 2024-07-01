package grpc_server

import (
	"context"
	"github.com/brianvoe/gofakeit"
	userdesc "github.com/neracastle/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

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
		Id: req.GetId(),
		User: &userdesc.User{
			Name:  gofakeit.BeerName(),
			Email: gofakeit.Email(),
			Role:  userdesc.Role_USER,
		},
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

func (s *Server) Update(ctx context.Context, req *userdesc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("called Update method with req: %v", req)

	return &emptypb.Empty{}, nil
}

func (s *Server) Delete(ctx context.Context, req *userdesc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("called Delete method with req: %v", req)

	return &emptypb.Empty{}, nil
}
