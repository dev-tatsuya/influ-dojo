package input

import (
	"influ-dojo/api/domain/repository"
	"log"
)

type RankRecord struct {
	DailyWorkRepo   repository.Work   `json:"-"`
	DailyResultRepo repository.Result `json:"-"`
}

func (in *RankRecord) RankDaily() error {
	works, err := in.DailyWorkRepo.LoadOrderByRanking()
	if err != nil {
		return err
	}

	log.Printf("works len %d: %+v", len(works), works)

	ranking := 1
	for i, work := range works {
		if i == 0 {
			work.MakeRankingPast()
			work.Ranking = ranking
		} else {
			if work.Point == works[i-1].Point {
				work.MakeRankingPast()
				work.Ranking = ranking
			} else {
				work.MakeRankingPast()
				ranking = i + 1
				work.Ranking = ranking
			}
		}

		if err := in.DailyWorkRepo.Save(work); err != nil {
			return err
		}
	}

	results, err := in.DailyResultRepo.LoadOrderByRanking()
	if err != nil {
		return err
	}

	ranking = 1
	for i, result := range results {
		if i == 0 {
			result.MakeRankingPast()
			result.Ranking = ranking
		} else {
			if result.Point == works[i-1].Point {
				result.MakeRankingPast()
				result.Ranking = ranking
			} else {
				result.MakeRankingPast()
				ranking = i + 1
				result.Ranking = ranking
			}
		}

		if err := in.DailyResultRepo.Save(result); err != nil {
			return err
		}
	}

	return nil
}
