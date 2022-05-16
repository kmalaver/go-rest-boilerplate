package user

import (
	"github.com/labstack/echo/v4"
)

func App(group *echo.Group) {
	repo := newUserRepository()
	newHandler(repo).routes(group)
}
