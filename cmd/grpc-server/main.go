package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/neracastle/auth/internal/config"
	usersrv "github.com/neracastle/auth/internal/grpc-server"
	userdesc "github.com/neracastle/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {

	cfg := config.MustLoad()

	conn, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))

	if err != nil {
		log.Fatal(color.RedString("failed to serve grpc server: %v", err))
	}

	log.Print(color.GreenString("UserAPI grpc server listening on: %s", conn.Addr().String()))

	gsrv := grpc.NewServer()
	reflection.Register(gsrv)

	userdesc.RegisterUserV1Server(gsrv, &usersrv.Server{})

	if err = gsrv.Serve(conn); err != nil {
		log.Fatal(color.RedString("failed to serve grpc server: %v", err))
	}
}
