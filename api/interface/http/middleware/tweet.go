package middleware

import (
	"github.com/labstack/echo/v4"
)

func MakeTweetHandlerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return next
	}
}
