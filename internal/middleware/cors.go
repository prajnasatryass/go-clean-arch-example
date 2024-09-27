package middleware

import (
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"net/http"
)

var corsConfig = echoMiddleware.CORSConfig{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete},
}
