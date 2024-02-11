package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net"

	descUser "github.com/a1exCross/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	grpcPort = 50051
)

var user *descUser.User

type server struct {
	descUser.UnimplementedUserV1Server
}

func (s server) Create(_ context.Context, req *descUser.CreateRequest) (*descUser.CreateResponse, error) {
	time := timestamppb.Now()

	id, err := rand.Int(rand.Reader, big.NewInt(123))
	if err != nil {
		return &descUser.CreateResponse{}, fmt.Errorf("failed to generate id: %v", err)
	}

	user = &descUser.User{
		Id:        id.Int64(),
		Info:      req.Info,
		CreatedAt: time,
		UpdatedAt: time,
	}

	log.Printf("user was created: %v", user)

	return &descUser.CreateResponse{
		Id: user.Id,
	}, nil
}

func (s server) Get(_ context.Context, req *descUser.GetRequest) (*descUser.GetResponse, error) {
	if req.Id == user.Id {
		log.Printf("user was find with id: %d", user.Id)
		return &descUser.GetResponse{
			User: user,
		}, nil
	}

	return &descUser.GetResponse{}, fmt.Errorf("user not found")
}

func (s server) Update(_ context.Context, req *descUser.UpdateRequest) (*emptypb.Empty, error) {
	time := timestamppb.Now()

	user.Info = &descUser.UserInfo{
		Name:  req.Name.Value,
		Email: req.Email.Value,
		Role:  req.Role,
	}

	user.UpdatedAt = time

	log.Printf("user was updated: %v", user)

	return &emptypb.Empty{}, nil
}

func (s server) Delete(_ context.Context, req *descUser.DeleteRequest) (*emptypb.Empty, error) {
	if req.Id == user.Id {
		log.Printf("user was deleted with id: %d", user.Id)

		user = nil
		return &emptypb.Empty{}, nil
	}

	return &emptypb.Empty{}, fmt.Errorf("user not found")
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("error at listen tcp: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	descUser.RegisterUserV1Server(s, server{})

	log.Printf("server listening at: %d", grpcPort)

	if err = s.Serve(lis); err != nil {
		if err != nil {
			log.Fatalf("error at grpc serve: %v", err)
		}
	}
}
