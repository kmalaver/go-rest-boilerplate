package rest

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type StructValidator interface {
	Validate() error
}

type StructSanitizer interface {
	Sanitize()
}

type CustomBinder struct{}

func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
	fmt.Println("Hooooooooooooooooooooooooooooo")

	// You may use default binder
	db := new(echo.DefaultBinder)
	if err := db.Bind(i, c); err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// Sanitize
	if v, ok := i.(StructSanitizer); ok {
		v.Sanitize()
	}

	// Validate
	switch v := i.(type) {
	case StructValidator: // validate with custom validator
		if err := v.Validate(); err != nil {
			return err
		}
	default: // validate with default validator
		if err := c.Validate(i); err != nil {
			fmt.Println("Error:", err)
			return err
		}
	}

	return nil
}

func Wrap[Req any](handler func(c echo.Context, req Req) error) func(c echo.Context) error {
	return func(c echo.Context) error {
		var req Req
		if err := c.Bind(&req); err != nil {
			return err
		}
		return handler(c, req)
	}
}
