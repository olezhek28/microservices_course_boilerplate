package main

import (
	pb "auth/pkg/user_v1"
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedUserV1Server
}

func (s *server) Create(_ context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	log.Println("create user request: ", req)
	return nil, nil
}

func (s *server) Get(_ context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	log.Println("get user request: ", req)
	return nil, nil
}

func (s *server) Update(_ context.Context, req *pb.UpdateRequest) (*emptypb.Empty, error) {
	log.Println("update user request: ", req)
	return nil, nil
}

func (s *server) Delete(_ context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error) {
	log.Println("delete user request: ", req)
	return nil, nil
}
