package auth

import (
	"rest/pkg/shared/models"

	"github.com/labstack/echo/v4"
)

func Roles(roles ...models.Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	}
}
