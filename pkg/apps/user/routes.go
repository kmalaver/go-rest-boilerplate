package user

import (
	"rest/pkg/shared/middlewares/auth"
	"rest/pkg/shared/models"

	"github.com/labstack/echo/v4"
)

func (h *handler) routes(api *echo.Group) {
	api.Use(
		auth.Apikey(),
	)

	api.GET("/", h.List)
	api.POST("/", h.Create, auth.Roles(models.RoleAdmin))
	api.GET("/:id", h.Get)
	api.PUT("/:id", h.Update, auth.Roles(models.RoleAdmin))
	api.DELETE("/:id", h.Delete, auth.Roles(models.RoleAdmin))
}
