package router

import (
	"github.com/andresxlp/api-mvc-arch/controllers"
	"github.com/andresxlp/api-mvc-arch/router/groups"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	server    *echo.Echo
	userGroup groups.UserGroup
}

func NewRouter(server *echo.Echo, userGroup groups.UserGroup) *Router {
	return &Router{
		server,
		userGroup,
	}
}

func (rtr *Router) Init() {
	rtr.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n",
	}))

	rtr.server.Use(middleware.Recover())

	basePath := rtr.server.Group("/api")

	basePath.GET("/health", controllers.HealthCheck)
	rtr.userGroup.Resource(basePath)

}
