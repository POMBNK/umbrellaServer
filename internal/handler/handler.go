package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service interface {
	DaysUntilCounter() string
}

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Status(c echo.Context) error {
	resp := h.service.DaysUntilCounter()
	err := c.String(http.StatusOK, resp)
	if err != nil {
		return err
	}
	return nil
}
