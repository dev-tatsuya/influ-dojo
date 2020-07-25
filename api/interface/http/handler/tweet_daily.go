package handler

import (
	"influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/input"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MakeDailyTweetHandler(
	bot client.Bot,
	work repository.DailyWork,
	result repository.DailyResult,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		in := input.DailyTweet{
			Bot:             bot,
			DailyWorkRepo:   work,
			DailyResultRepo: result,
		}

		if err := in.TweetDaily(); err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}
