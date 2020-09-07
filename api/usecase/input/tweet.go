package input

import (
	"influ-dojo/api/domain/client"
	queryService "influ-dojo/api/usecase/query"
)

type Tweet struct {
	Path    string
	Bot     client.Bot
	Ranking queryService.Ranking
}

func (in *Tweet) Tweet() error {
	top3, err := in.Ranking.LoadRankingTop3()
	if err != nil {
		return err
	}

	return in.Bot.Tweet(top3, in.Path)
}
