package middlewares

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func Log() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			res := c.Response()
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()
			l := log.Info().
				Str("method", req.Method).
				Str("path", c.Path()).
				Int("status", res.Status).
				Str("duration", stop.Sub(start).String())
			if err != nil {
				l.Err(err)
			}
			l.Send()
			return err
		}
	}
}
