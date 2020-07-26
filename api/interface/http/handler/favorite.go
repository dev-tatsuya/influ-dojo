package handler

import (
	"influ-dojo/api/domain/client"
	"influ-dojo/api/usecase/input"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MakeFavoriteHandler(bot client.Bot) echo.HandlerFunc {
	return func(c echo.Context) error {
		in := &input.Favorite{Bot: bot}

		if err := in.Favorite(); err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}
