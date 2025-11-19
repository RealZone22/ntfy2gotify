package api

import (
	"ntfy2gotify/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoggerMiddleware() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		BeforeNextFunc: func(c echo.Context) {
			c.Set("customValueFromContext", 42)
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			utils.Logger.Debug().Str("uri", v.URI).Int("status", v.Status).
				Str("ip", c.RealIP()).Str("method", c.Request().Method).Msg("Incoming request")
			return nil
		},
	})
}
