package input

import (
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/dto"
)

type RankingAll struct {
	RankingRepo repository.Ranking
}

func (in *RankingAll) GetRankingAll() (*dto.RankingAll, error) {
	all, err := in.RankingRepo.LoadAll()
	if err != nil {
		return nil, err
	}

	return all, nil
}
