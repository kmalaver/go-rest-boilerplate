package rest

import (
	"rest/pkg/shared/rest"

	"github.com/labstack/echo/v4"
)

func (h *handler) Routes(api *echo.Group) {
	api.GET("/", h.ListCompanies).Name = "companies.listCompanies"
	api.POST("/", h.CreateCompany).Name = "companies.createCompany"
	api.PUT("/:id/", rest.Wrap(h.UpdateCompany)).Name = "companies.updateCompany"
}
