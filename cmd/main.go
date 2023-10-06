package main

import (
	"context"
	desc "github.com/MikhailRibalkov/chat-server/pkg/chatServer_v1/pkg/chatServer_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

// Написать тесты!
const grpcPort = ":8082"

type chatServer struct {
	desc.UnimplementedChatServerV1Server
}

func (s *chatServer) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Got message from: %s\nmessage: %s\n", req.From, req.Text)

	return &emptypb.Empty{}, nil
}

func (s *chatServer) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Got message usernames: %v", req.GetUsernames())

	return &desc.CreateResponse{}, nil
}

func (s *chatServer) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Request for delete user with id: %d", req.GetId())

	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatServerV1Server(s, &chatServer{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
