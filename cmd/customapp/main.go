package main

import (
	"context"
	"github.com/POMBNK/umbrellaServer/internal/handler"
	"github.com/POMBNK/umbrellaServer/internal/service"
	"github.com/POMBNK/umbrellaServer/pkg/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const port = "8080"

func main() {

	s := service.New()
	h := handler.New(s)
	srv := new(server.Server)
	go func() {
		if err := srv.Run(port, h.InitRoutes()); err != nil {
			log.Fatalf("error while running server %s", err)
		}
	}()

	log.Println("Service started")

	// Graceful Shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Service shutting down...")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occurred on server shutting down: %s", err.Error())
	}
}
