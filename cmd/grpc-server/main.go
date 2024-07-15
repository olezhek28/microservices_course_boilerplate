package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/fatih/color"
	"github.com/jackc/pgx/v5"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/neracastle/auth/internal/config"
	usersrv "github.com/neracastle/auth/internal/grpc-server"
	userdesc "github.com/neracastle/auth/pkg/user_v1"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	//нижележащее в дальнейшем вынесем в di и слои
	logger := setupLogger(cfg.Env)

	logger.Debug("connecting to postgres",
		slog.String("Host", cfg.Postgres.Host),
		slog.Int("Port", cfg.Postgres.Port))

	ctx := context.Background()
	connect, err := pgx.Connect(ctx, cfg.Postgres.DSN())
	if err != nil {
		log.Fatalf("failed to connect to pg: %v", err)
	}
	defer connect.Close(ctx)

	err = connect.Ping(ctx)
	if err != nil {
		log.Fatalf("failed ping to pg: %v", err)
	}

	logger.Debug("connected to postgres",
		slog.String("Host", cfg.Postgres.Host),
		slog.Int("Port", cfg.Postgres.Port))

	logger.Debug("starting grpc server",
		slog.String("Host", cfg.GRPC.Host),
		slog.Int("Port", cfg.GRPC.Port))

	conn, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.GRPC.Host, cfg.GRPC.Port))
	if err != nil {
		log.Fatal(color.RedString("failed to serve grpc server: %v", err))
	}

	gsrv := grpc.NewServer()
	reflection.Register(gsrv)

	userdesc.RegisterUserV1Server(gsrv, usersrv.NewServer(logger, connect))

	logger.Info("UserAPI service started", slog.String("Host", cfg.GRPC.Host), slog.Int("Port", cfg.GRPC.Port))

	if err = gsrv.Serve(conn); err != nil {
		log.Fatal(color.RedString("failed to serve grpc server: %v", err))
	}
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		break
	case envDev:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		break
	case envProd:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}
