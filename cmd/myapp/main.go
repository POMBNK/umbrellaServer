package main

import (
	"context"
	"github.com/POMBNK/umbrellaServer/internal/handler"
	"github.com/POMBNK/umbrellaServer/internal/service"
	"log"
	"os"
	"os/signal"
	"time"
)

const port = ":8080"

func main() {

	s := service.New()
	h := handler.New(s)
	// Echo has method Start() to run http server in a handler domain. So, if you want, you can create a custom server from
	// "net/http" using a *http.Server. Echo pkg also use it under hood.
	httpserver := h.InitRoutes()

	go func() {
		if err := httpserver.Start(port); err != nil {
			log.Fatalf("Server start error: %s", err)
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := httpserver.Shutdown(ctx); err != nil {
		httpserver.Logger.Fatal(err)
	}
}
