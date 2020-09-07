//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=../../mock/$GOPACKAGE/mock_$GOFILE

package queryService

import "influ-dojo/api/usecase/dto"

type Ranking interface {
	LoadRankingAll() (*dto.RankingAll, error)
	LoadRankingTop3() (*dto.Top3, error)
}
