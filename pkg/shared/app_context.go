package shared

import (
	"github.com/labstack/echo/v4"
)

type App func(*echo.Group)
type Apps []App
type Middlewares []echo.MiddlewareFunc

func Route(path string, app App) App {
	return func(group *echo.Group) {
		routeGroup := group.Group(path)
		app(routeGroup)
	}
}
