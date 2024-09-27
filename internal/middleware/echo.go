package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/prajnasatryass/go-clean-arch-example/config"
)

func NewEcho(config config.Config) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Use(echoMiddleware.CORSWithConfig(corsConfig))
	e.Validator = &RequestValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
	e.Debug = config.Server.Debug
	return e
}
