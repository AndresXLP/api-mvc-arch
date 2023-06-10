package controllers

import (
	"net/http"

	"github.com/andresxlp/api-mvc-arch/dto"
	"github.com/andresxlp/api-mvc-arch/services"
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	CreateUser(c echo.Context) error
}

type userHandler struct {
	svc services.UserServices
}

func NewUserController(svc services.UserServices) UserHandler {
	return &userHandler{
		svc,
	}
}

func (hand *userHandler) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()
	user := dto.User{}

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "No llego un string")
	}

	if err := user.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := hand.svc.CreateUser(ctx, user)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, "User Created Successfully")
}
