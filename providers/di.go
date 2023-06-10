package providers

import (
	"github.com/andresxlp/api-mvc-arch/controllers"
	"github.com/andresxlp/api-mvc-arch/db"
	"github.com/andresxlp/api-mvc-arch/db/repo"
	"github.com/andresxlp/api-mvc-arch/router"
	"github.com/andresxlp/api-mvc-arch/router/groups"
	"github.com/andresxlp/api-mvc-arch/services"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()

	_ = Container.Provide(func() *echo.Echo {
		return echo.New()
	})

	_ = Container.Provide(db.NewConnection)

	_ = Container.Provide(router.NewRouter)

	_ = Container.Provide(groups.NewUserRouter)

	_ = Container.Provide(controllers.NewUserController)

	_ = Container.Provide(repo.NewRepo)

	_ = Container.Provide(services.NewUserService)

	return Container
}
