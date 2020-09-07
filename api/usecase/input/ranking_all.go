package input

import (
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/dto"
)

type RankingAll struct {
	RankingRepo repository.Ranking
}

func (in *RankingAll) GetRankingAll() (*dto.RankingAll, error) {
	return in.RankingRepo.LoadAll()
}
