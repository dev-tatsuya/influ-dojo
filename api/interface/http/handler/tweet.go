package handler

import (
	"influ-dojo/api/domain/client"
	"influ-dojo/api/usecase/input"
	queryService "influ-dojo/api/usecase/query"
	"log"
	"net/http"
	"path"

	"github.com/labstack/echo/v4"
)

func MakeTweetHandler(
	bot client.Bot,
	ranking queryService.Ranking,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		in := input.Tweet{
			Path:    path.Base(c.Path()),
			Bot:     bot,
			Ranking: ranking,
		}

		if err := in.Tweet(); err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}

func Tweet(
	bot client.Bot,
	ranking queryService.Ranking,
	path string,
) {
	in := input.Tweet{
		Path:    path,
		Bot:     bot,
		Ranking: ranking,
	}

	if err := in.Tweet(); err != nil {
		log.Printf("%+v", err)
	}
}
