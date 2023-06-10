package main

import (
	"fmt"
	"log"

	"github.com/andresxlp/api-mvc-arch/config"
	"github.com/andresxlp/api-mvc-arch/providers"
	"github.com/andresxlp/api-mvc-arch/router"
	"github.com/labstack/echo/v4"
)

func main() {
	container := providers.BuildContainer()
	port := config.Environments().Port

	if err := container.Invoke(func(server *echo.Echo, router *router.Router) {
		router.Init()
		server.Logger.Fatal(server.Start(fmt.Sprintf(":%d", port)))
	}); err != nil {
		log.Fatal(err)
	}
}
