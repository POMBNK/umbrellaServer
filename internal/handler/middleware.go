package handler

import (
	"github.com/labstack/echo/v4"
	"log"
	"strings"
)

const (
	adminRole      = "admin"
	userRoleHeader = "User-Role"
)

func (h *Handler) RoleAuth(next echo.HandlerFunc) echo.HandlerFunc {
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
