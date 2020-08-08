package handler

import (
	"influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/input"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MakeClassifyTweetsHandler(
	client client.Tweet,
	dailyWork repository.Work,
	weeklyWork repository.Work,
	monthlyWork repository.Work,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		in := &input.ClassifyTweets{
			TweetClient:     client,
			DailyWorkRepo:   dailyWork,
			WeeklyWorkRepo:  weeklyWork,
			MonthlyWorkRepo: monthlyWork,
		}

		if err := in.Classify(); err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}

func ClassifyDailyTweets(
	client client.Tweet,
	dailyWork repository.Work,
	weeklyWork repository.Work,
	monthlyWork repository.Work,
) {
	in := &input.ClassifyTweets{
		TweetClient:     client,
		DailyWorkRepo:   dailyWork,
		WeeklyWorkRepo:  weeklyWork,
		MonthlyWorkRepo: monthlyWork,
	}

	if err := in.Classify(); err != nil {
		log.Printf("%+v", err)
	}
}
