package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Health struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func HealthCheck(c echo.Context) error {
	response := &Health{
		Code:    http.StatusOK,
		Message: "Active!",
	}
	return c.JSON(http.StatusOK, response)
}
