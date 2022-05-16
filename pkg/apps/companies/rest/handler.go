package rest

import "github.com/labstack/echo/v4"

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

type DetailPathParams struct {
	Id int64 `path:"id"`
}

func (*handler) CreateCompany(c echo.Context) error {
	type CreateCompanyBody struct {
		Name    string `json:"name"`
		Address string `json:"addresss"`
	}
	var body CreateCompanyBody
	if err := c.Bind(&body); err != nil {
		return err
	}

	return c.JSON(200, "")
}

func (*handler) ListCompanies(c echo.Context) error {
	type ListCompaniesParams struct {
		Search *string `query:"search"`
		Limit  *int    `query:"limit"`
		Offset *int    `query:"offset"`
	}
	var params ListCompaniesParams
	if err := c.Bind(&params); err != nil {
		return err
	}
	// res, err := h.service.ListCompanies(params)
	// if err != nil {
	// 	return err
	// }
	// return c.JSON(200, res)

	return c.JSON(200, params)
}

type UpdateConpanyRequest struct {
	Id   int64  `json:"id" param:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"min=18"`
}

func (*handler) UpdateCompany(c echo.Context, req UpdateConpanyRequest) error {
	return c.JSON(200, req)
}
