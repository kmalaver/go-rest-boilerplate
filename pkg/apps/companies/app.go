package companies

import (
	"rest/pkg/apps/companies/rest"

	"github.com/labstack/echo/v4"
)

func App(group *echo.Group) {
	rest.NewHandler().Routes(group)
}
