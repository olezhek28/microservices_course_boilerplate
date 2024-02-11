package main

import (
	"context"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	desc "github.com/mchekalov/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	desc.UnimplementedUserV1Server
}

func (s *server) Create(ctx context.Context, in *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create User %v", in.GetInfo())
	return &desc.CreateResponse{
		Id: 2,
	}, nil
}

func (s *server) Get(ctx context.Context, in *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("User ID %v", in.GetId())
	return &desc.GetResponse{
		User: &desc.User{
			Id: in.GetId(),
			Info: &desc.UserInfo{
				Name:  gofakeit.Name(),
				Email: gofakeit.Email(),
				Role:  0,
			},
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

func (s *server) Update(ctx context.Context, in *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("Update user %v", in.GetWrap())
	return new(emptypb.Empty), nil
}

func (s *server) Delete(ctx context.Context, in *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Delete user %v", in.GetId())
	return new(emptypb.Empty), nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("fatal to listen %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve; %v", err)
	}
}
