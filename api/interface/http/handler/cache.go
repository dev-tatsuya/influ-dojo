package handler

import (
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/input"
	queryService "influ-dojo/api/usecase/query"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MakeCacheHandler(
	rankingQuery queryService.Ranking,
	rankingRepo repository.Ranking,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		in := &input.Cache{
			RankingQuery: rankingQuery,
			RankingRepo:  rankingRepo,
		}

		if err := in.Cache(); err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}
