package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/algol-84/auth/pkg/auth_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedAuthV1Server
}

// Create User
func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create new user: %+v with role %v", req.Info, req.Info.Role)

	return &desc.CreateResponse{
		Id: 0,
	}, nil
}

// Get User info by ID
func (s *server) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Get user by ID: %d", req.GetId())

	return &desc.GetResponse{
		User: &desc.GetUser{
			Id:        req.GetId(),
			Name:      "Noname",
			Email:     "empty@mail.ru",
			Role:      desc.Role_USER,
			CreatedAt: timestamppb.New(time.Now()),
			UpdatedAt: timestamppb.New(time.Now()),
		},
	}, nil
}

// Update User info
func (s *server) Update(_ context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("Update user by ID: %d", req.Id)

	return &emptypb.Empty{}, nil
}

// Delete User by ID
func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Delete user by ID: %d", req.Id)

	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterAuthV1Server(s, &server{})

	log.Printf("auth server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
