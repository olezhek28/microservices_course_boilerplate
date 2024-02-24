package main

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/dmtrybogdanov/auth/pkg/user_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
)

const grpcPort = 50051

type server struct {
	user_v1.UnimplementedUserV1Server
}

func (s *server) Create(ctx context.Context, req *user_v1.CreateRequest) (*user_v1.CreateResponse, error) {
	log.Printf("Received CreateRequest: %v", req)

	return &user_v1.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Get(ctx context.Context, req *user_v1.GetRequest) (*user_v1.GetResponse, error) {
	log.Printf("Received GetRequest: %v", req)

	return &user_v1.GetResponse{
		Id:        req.Id,
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      user_v1.Role_USER,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

func (s *server) Update(ctx context.Context, in *user_v1.UpdateRequest) (*user_v1.UpdateResponse, error) {
	log.Printf("Received UpdateRequest: %v", in)
	return &user_v1.UpdateResponse{ResponseUpdate: &empty.Empty{}}, nil
}

func (s *server) Delete(ctx context.Context, in *user_v1.DeleteRequest) (*user_v1.DeleteResponse, error) {
	log.Printf("Received DeleteRequest: %v", in)
	return &user_v1.DeleteResponse{ResponseDelete: &empty.Empty{}}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	user_v1.RegisterUserV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
