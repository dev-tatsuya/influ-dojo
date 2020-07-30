package handler

import (
	"influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/input"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MakeRecordMonthlyHandler(
	follower client.Follower,
	user repository.User,
	work repository.Work,
	result repository.Result,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		in := input.MonthlyRecord{
			FollowerClient:    follower,
			UserRepo:          user,
			MonthlyWorkRepo:   work,
			MonthlyResultRepo: result,
		}

		if err := in.RecordMonthly(); err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}
