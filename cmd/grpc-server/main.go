package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/fatih/color"

	"github.com/neracastle/auth/internal/app"
)

func main() {

	ap := app.NewApp(context.Background())

	go func() {
		err := ap.Start()
		if err != nil {
			log.Fatal(color.RedString("failed to start app: %v", err))
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("received signal, graceful shutdown", sig)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	ap.Shutdown(ctx)
}
