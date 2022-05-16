package server

import (
	"fmt"
	"rest/config"
	"rest/pkg/shared/rest"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewApp() *echo.Echo {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true
	e.Validator = rest.Validator
	e.Binder = &rest.CustomBinder{}
	// e.HTTPErrorHandler = rest.CustomHTTPErrorHandler
	e.Pre(middleware.AddTrailingSlash())

	group := e.Group("api")

	for _, middleware := range middlewares {
		group.Use(middleware)
	}

	for _, app := range apps {
		app(group)
	}

	e.Static("/static", "public")

	//fmt.Println(e.Reverse("companies.updateCompany", 45))
	return e
}

func RunServer() {
	app := NewApp()
	Info(app)
	app.Logger.Fatal(app.Start(config.Server.GetUrl()))
}

func Info(e *echo.Echo) {
	url := "http://" + config.Server.GetUrl()
	fmt.Println()
	fmt.Println("🚀 Server running on:", color.GreenString(url))
	fmt.Println()
}
