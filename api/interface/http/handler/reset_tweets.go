package handler

import (
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/input"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MakeResetTweetsHandler(
	work repository.Work,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		in := &input.ResetTweets{
			WorkRepo: work,
		}

		if err := in.ResetTweets(); err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}

func ResetTweetsCount(
	work repository.Work,
) {
	in := &input.ResetTweets{
		WorkRepo: work,
	}

	if err := in.ResetTweets(); err != nil {
		log.Printf("%+v", err)
	}
}
