package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Greet struct {
	Message string `json:"message"`
}

func MakeHelloHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, &Greet{Message: "Hello, Flutter Web"})
	}
}
