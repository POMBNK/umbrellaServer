package main

import (
	"github.com/POMBNK/umbrellaServer/internal/handler"
	"github.com/POMBNK/umbrellaServer/internal/service"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {

	s := service.New()
	h := handler.New(s)

	//Todo: mb better encapsulate echo instance bellow
	server := echo.New()
	server.GET("/button", h.Status, h.RoleAuth)
	if err := server.Start(":8080"); err != nil {
		log.Fatalf("Server start error %s", err)
	}
}
