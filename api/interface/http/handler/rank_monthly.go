package handler

import (
	"influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/input"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MakeMonthlyRankHandler(
	follower client.Follower,
	user repository.User,
	work repository.MonthlyWork,
	result repository.MonthlyResult,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		in := input.MonthlyRank{
			FollowerClient:    follower,
			UserRepo:          user,
			MonthlyWorkRepo:   work,
			MonthlyResultRepo: result,
		}

		out, err := in.GetMonthlyRank()
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, out)
	}
}
