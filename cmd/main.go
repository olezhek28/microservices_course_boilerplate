package main

import (
	"context"
	"log"
	"net"

	"github.com/MikhailRibalkov/chat-server/pkg/chatServer_v1/pkg"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const grpcPort = ":8082"

type chatServer struct {
	chatServer_v1.UnimplementedChatServerV1Server
}

func (s *chatServer) SendMessage(ctx context.Context, req *chatServer_v1.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Got message from: %s\nmessage: %s\n", req.From, req.Text)

	return &emptypb.Empty{}, nil
}

func (s *chatServer) Create(ctx context.Context, req *chatServer_v1.CreateRequest) (*chatServer_v1.CreateResponse, error) {
	log.Printf("Got message usernames: %v", req.GetUsernames())

	return &chatServer_v1.CreateResponse{}, nil
}

func (s *chatServer) Delete(ctx context.Context, req *chatServer_v1.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Request for delete user with id: %d", req.GetId())

	return &emptypb.Empty{}, nil
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	chatServer_v1.RegisterChatServerV1Server(s, &chatServer{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
