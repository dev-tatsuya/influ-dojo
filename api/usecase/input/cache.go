package input

import (
	"influ-dojo/api/domain/repository"
	queryService "influ-dojo/api/usecase/query"
)

type Cache struct {
	RankingQuery queryService.Ranking
	RankingRepo  repository.Ranking
}

func (in *Cache) Cache() error {
	all, err := in.RankingQuery.LoadRankingAll()
	if err != nil {
		return err
	}

	return in.RankingRepo.Store(all)
}
