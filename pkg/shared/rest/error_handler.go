package rest

import (
	"fmt"
	"net/http"
	"rest/pkg/shared/errors"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Error *errors.Error `json:"error"`
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	switch e := err.(type) {
	case *echo.HTTPError:
		c.JSON(
			e.Code,
			ErrorResponse{
				Error: errors.ErrDefaultHTTPError(
					fmt.Sprintf("%v", e.Message),
					e.Code,
				),
			},
		)
	case *errors.Error:
		c.JSON(
			e.Status,
			ErrorResponse{
				Error: e,
			},
		)
	default:
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: errors.ErrDefaultHTTPError(e.Error()),
		})
	}
}
