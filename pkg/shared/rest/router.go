package rest

import "github.com/labstack/echo/v4"

type GetHandler interface {
	Get(c echo.Context) error
}

type ListHandler interface {
	List(c echo.Context) error
}

type CreateHandler interface {
	Create(c echo.Context) error
}

type UpdateHandler interface {
	Update(c echo.Context) error
}

type DeleteHandler interface {
	Delete(c echo.Context) error
}

func DefaultRouter(api *echo.Group, h interface{}) {
	if h, ok := h.(GetHandler); ok {
		api.GET("/:id", h.Get)
	}
	if h, ok := h.(ListHandler); ok {
		api.GET("/", h.List)
	}
	if h, ok := h.(CreateHandler); ok {
		api.POST("/", h.Create)
	}
	if h, ok := h.(UpdateHandler); ok {
		api.PUT("/:id", h.Update)
	}
	if h, ok := h.(DeleteHandler); ok {
		api.DELETE("/:id", h.Delete)
	}
}
