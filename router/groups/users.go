package groups

import (
	"github.com/andresxlp/api-mvc-arch/controllers"
	"github.com/labstack/echo/v4"
)

type UserGroup interface {
	Resource(g *echo.Group)
}

type userRouter struct {
	userHand controllers.UserHandler
}

func NewUserRouter(userHand controllers.UserHandler) UserGroup {
	return &userRouter{
		userHand,
	}
}

func (group userRouter) Resource(g *echo.Group) {
	groupPath := g.Group("/user")

	groupPath.POST("", group.userHand.CreateUser)
}
