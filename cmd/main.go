package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	adminRole      = "admin"
	userRoleHeader = "User-Role"
)

func main() {

	server := echo.New()
	//Todo: make middleware for admin role auth

	server.GET("/button", Handler, Middleware)

	if err := server.Start(":8080"); err != nil {
		log.Fatalf("Server start error %s", err)
	}
}

func Handler(c echo.Context) error {

	endTime := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	dur := time.Until(endTime)
	s := fmt.Sprintf("Days until 25 JAN 2025:%d", hoursToDays(dur))
	err := c.String(http.StatusOK, s)
	if err != nil {
		return err
	}
	return nil
}

func hoursToDays(d time.Duration) int64 {
	return int64(d.Hours() / 24)
}

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Request().Header.Get(userRoleHeader)

		if strings.EqualFold(role, adminRole) {
			log.Println("red button user detected!")
		} else {
			log.Println("access denied")
		}

		if err := next(c); err != nil {
			return err
		}
		return nil
	}
}
