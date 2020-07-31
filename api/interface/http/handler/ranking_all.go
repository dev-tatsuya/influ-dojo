package handler

import (
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/input"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MakeRankingAllHandler(
	rankingRepo repository.Ranking,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		in := &input.RankingAll{
			RankingRepo: rankingRepo,
		}

		out, err := in.GetRankingAll()
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, out)
	}
}
