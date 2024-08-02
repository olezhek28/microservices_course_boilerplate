package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fatih/color"

	"github.com/neracastle/auth/internal/app"
)

func main() {

	ctx := context.Background()
	ap := app.NewApp(ctx)

	go func() {
		err := ap.Start()
		if err != nil {
			log.Fatal(color.RedString("failed to start app: %v", err))
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan
	log.Println("received signal, graceful shutdown", sig)
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	ap.Shutdown(ctx)
}
