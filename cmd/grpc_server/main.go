package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/brianvoe/gofakeit/v6"
	desc "github.com/sarastee/auth/pkg/user_api_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserAPIV1Server
}

var testUser *desc.User

// Create User
func (s *server) Create(_ context.Context, _ *desc.CreateRequest) (*desc.CreateResponse, error) {

	testUser = &desc.User{
		Id: 0,
		Info: &desc.UserInfo{
			Name:  gofakeit.Name(),
			Email: gofakeit.Email(),
			Role:  desc.Role_user,
		},
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}

	log.Printf("User has been created: %v", testUser)

	return &desc.CreateResponse{
		Id: testUser.Id,
	}, nil

	// Create Request error
}

// Get User info
func (s *server) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {

	if req.Id != testUser.Id {
		log.Printf("User with id: %v not found", req.Id)
		return nil, fmt.Errorf("user not found")
	}
	log.Printf("User id: %d was found", req.GetId())

	return &desc.GetResponse{
		User: testUser,
	}, nil

	// Get Request error
}

// Update User
func (s *server) Update(_ context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {

	testUser.Info = &desc.UserInfo{
		Name:  req.Update.Name.Value,
		Email: req.Update.Email.Value,
		Role:  req.Update.Role,
	}

	testUser.UpdatedAt = timestamppb.Now()

	log.Printf("User with ID: %d has been updated", req.Id)
	return &emptypb.Empty{}, nil
}

// Delete User
func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	if req.Id != testUser.Id {
		log.Printf("User with id: %v not found", req.Id)
		return nil, fmt.Errorf("user not found")
	}
	log.Printf("User id: %d was found", req.Id)

	testUser = nil
	log.Printf("User id: %d has been deleted", req.Id)

	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserAPIV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
