package server

import (
	"rest/pkg/shared"
	customMiddlewares "rest/pkg/shared/middlewares"

	"github.com/labstack/echo/v4/middleware"
)

var middlewares = shared.Middlewares{
	middleware.Recover(),
	customMiddlewares.Log(),
	middleware.CORS(),
	// middleware.CSRF(),
	middleware.Secure(),
}
