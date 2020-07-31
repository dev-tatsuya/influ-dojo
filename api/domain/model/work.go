package model

import "influ-dojo/api/domain/utils"

type Work struct {
	ScreenName             string
	TweetsCount            int
	IncreaseTweetsCount    int
	FavoritesCount         int
	IncreaseFavoritesCount int
	Point                  int
	Ranking                int
	LastRanking            int
}

func (work *Work) MakeRankingPast() {
	work.LastRanking = work.Ranking
}

func (work *Work) UpdateCount(latestTweetsCount, latestFavoritesCount int) {
	work.IncreaseTweetsCount = utils.Sub(latestTweetsCount, work.TweetsCount)
	work.IncreaseFavoritesCount = utils.Sub(latestFavoritesCount, work.FavoritesCount)
	work.setPoint()
	work.TweetsCount = latestTweetsCount
	work.FavoritesCount = latestFavoritesCount
}

func (work *Work) setPoint() {
	work.Point = work.IncreaseTweetsCount*200 + work.IncreaseFavoritesCount
}
