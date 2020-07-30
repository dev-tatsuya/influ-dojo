package handler

import (
	"influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/input"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MakeRecordWeeklyHandler(
	follower client.Follower,
	user repository.User,
	work repository.Work,
	result repository.Result,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		in := input.WeeklyRecord{
			FollowerClient:   follower,
			UserRepo:         user,
			WeeklyWorkRepo:   work,
			WeeklyResultRepo: result,
		}

		if err := in.RecordWeekly(); err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}
