package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func MakeHelloHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "hello!\n")
	}
}
