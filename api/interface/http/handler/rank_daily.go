package handler

import (
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/input"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MakeRankDailyHandler(
	work repository.Work,
	result repository.Result,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		in := input.RankRecord{
			DailyWorkRepo:   work,
			DailyResultRepo: result,
		}

		if err := in.RankDaily(); err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}
