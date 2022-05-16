package user

import (
	"rest/pkg/shared/models"
	"rest/pkg/shared/rest"

	"github.com/labstack/echo/v4"
)

type handler struct {
	repo *userRepository
}

func newHandler(repo *userRepository) *handler {
	return &handler{
		repo,
	}
}

func (h *handler) Get(c echo.Context) error {
	return c.String(200, "Hello, World!")
}

func (h *handler) List(c echo.Context) error {
	filters := rest.GetFilters(c)
	users, err := h.repo.List(filters.ToBuilder()) // TODO: change to builder here, filters should be part of domain
	if err != nil {
		return err
	}
	return c.JSON(200, users)
}

func (h *handler) Create(c echo.Context) error {
	var body models.User
	if err := c.Bind(&body); err != nil {
		return err
	}
	return c.JSON(201, body)
}

func (h *handler) Update(c echo.Context) error {
	return c.String(200, "Hello, World! list")
}

func (h *handler) Delete(c echo.Context) error {
	return c.String(200, "Hello, World! list")
}
