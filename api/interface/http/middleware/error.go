package middleware

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/multierr"
)

func MakeErrorHandlerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				for _, e := range multierr.Errors(err) {
					log.Printf("%+v\n", e)

					return c.String(http.StatusInternalServerError, "internal server error")
				}
			}

			return nil
		}
	}
}
