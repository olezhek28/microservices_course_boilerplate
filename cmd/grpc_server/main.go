package main

import (
	"fmt"
	"log"
	"net"

	descUser "github.com/a1exCross/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = 50051
)

type server struct {
	descUser.UnimplementedUserV1Server
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
