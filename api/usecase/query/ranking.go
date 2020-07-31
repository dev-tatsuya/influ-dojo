package queryService

import "influ-dojo/api/usecase/dto"

type Ranking interface {
	LoadRankingAll() (*dto.RankingAll, error)
}
