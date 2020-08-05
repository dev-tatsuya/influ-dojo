package handler

import (
	"influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/input"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MakeRecordUsersHandler(
	follower client.Follower,
	user repository.User,
	dailyWork repository.Work,
	dailyResult repository.Result,
	weeklyWork repository.Work,
	weeklyResult repository.Result,
	monthlyWork repository.Work,
	monthlyResult repository.Result,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		in := input.RecordUsers{
			FollowerClient:    follower,
			UserRepo:          user,
			DailyWorkRepo:     dailyWork,
			DailyResultRepo:   dailyResult,
			WeeklyWorkRepo:    weeklyWork,
			WeeklyResultRepo:  weeklyResult,
			MonthlyWorkRepo:   monthlyWork,
			MonthlyResultRepo: monthlyResult,
		}

		if err := in.RecordUsers(); err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}

func RecordNewUsers(
	follower client.Follower,
	user repository.User,
	dailyWork repository.Work,
	dailyResult repository.Result,
	weeklyWork repository.Work,
	weeklyResult repository.Result,
	monthlyWork repository.Work,
	monthlyResult repository.Result,
) {
	in := input.RecordUsers{
		FollowerClient:    follower,
		UserRepo:          user,
		DailyWorkRepo:     dailyWork,
		DailyResultRepo:   dailyResult,
		WeeklyWorkRepo:    weeklyWork,
		WeeklyResultRepo:  weeklyResult,
		MonthlyWorkRepo:   monthlyWork,
		MonthlyResultRepo: monthlyResult,
	}

	if err := in.RecordUsers(); err != nil {
		log.Printf("%+v", err)
	}
}
