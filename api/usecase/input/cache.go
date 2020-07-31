package input

import (
	"fmt"
	queryService "influ-dojo/api/usecase/query"
)

type Cache struct {
	RankingQuery queryService.Ranking
}

func (in *Cache) Cache() error {
	all, err := in.RankingQuery.LoadRankingAll()
	if err != nil {
		return err
	}

	// redisに格納
	fmt.Println(all)

	return nil
}
