package app

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/neracastle/go-libs/pkg/sys/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpc_server "github.com/neracastle/auth/internal/grpc-server"
	"github.com/neracastle/auth/pkg/user_v1"
)

// App приложение
type App struct {
	grpc        *grpc.Server
	srvProvider *serviceProvider
}

// NewApp новый экземпляр приложения
func NewApp(ctx context.Context) *App {
	app := &App{srvProvider: newServiceProvider()}
	app.init(ctx)
	return app
}

func (a *App) init(ctx context.Context) {
	lg := logger.SetupLogger(a.srvProvider.Config().Env)
	a.grpc = grpc.NewServer()

	reflection.Register(a.grpc)
	user_v1.RegisterUserV1Server(a.grpc, grpc_server.NewServer(lg, a.srvProvider.UsersService(ctx)))
}

// Start запускает сервис на прием запросов
func (a *App) Start() error {

	conn, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.srvProvider.Config().GRPC.Host, a.srvProvider.Config().GRPC.Port))
	if err != nil {
		return err
	}

	log.Printf("UserAPI service started on %s:%d\n", a.srvProvider.Config().GRPC.Host, a.srvProvider.Config().GRPC.Port)

	if err = a.grpc.Serve(conn); err != nil {
		return err
	}

	return nil
}

// Shutdown мягко закрывает все соединения и службы
func (a *App) Shutdown(ctx context.Context) {

	allClosed := make(chan struct{})
	go func() {
		_ = a.srvProvider.dbc.Close()
		a.grpc.GracefulStop()
		close(allClosed)
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case <-allClosed:
			return
		}
	}
}
