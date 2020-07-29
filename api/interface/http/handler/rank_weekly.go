package handler

import (
	"influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/input"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MakeWeeklyRankHandler(
	follower client.Follower,
	user repository.User,
	work repository.WeeklyWork,
	result repository.WeeklyResult,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		in := input.WeeklyRank{
			FollowerClient:   follower,
			UserRepo:         user,
			WeeklyWorkRepo:   work,
			WeeklyResultRepo: result,
		}

		out, err := in.GetWeeklyRank()
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, out)
	}
}
