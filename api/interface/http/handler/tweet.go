package handler

import (
	"influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/input"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MakeTweetHandler(
	bot client.Bot,
	work repository.Work,
	result repository.Result,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		in := input.Tweet{
			Path:       c.Path(),
			Bot:        bot,
			WorkRepo:   work,
			ResultRepo: result,
		}

		if err := in.Tweet(); err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}

func Tweet(
	bot client.Bot,
	work repository.Work,
	result repository.Result,
	path string,
) {
	in := input.Tweet{
		Path:       path,
		Bot:        bot,
		WorkRepo:   work,
		ResultRepo: result,
	}

	if err := in.Tweet(); err != nil {
		log.Printf("%+v", err)
	}
}
