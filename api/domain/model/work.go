package model

import (
	"influ-dojo/api/domain/utils"
	"math"
)

type Work struct {
	ScreenName             string
	TweetsCount            int
	IncreaseTweetsCount    int
	MyTweetsCount          int
	RepliesCount           int
	FavoritesCount         int
	IncreaseFavoritesCount int
	Point                  float64
	Ranking                int
	LastRanking            int
}

func (work *Work) MakeRankingPast() {
	work.LastRanking = work.Ranking
}

func (work *Work) UpdateCount(latestTweetsCount, latestFavoritesCount int) {
	work.IncreaseTweetsCount = utils.Sub(latestTweetsCount, work.TweetsCount)
	work.IncreaseFavoritesCount = utils.Sub(latestFavoritesCount, work.FavoritesCount)
	work.TweetsCount = latestTweetsCount
	work.FavoritesCount = latestFavoritesCount
}

func (work *Work) SetPoint() {
	myTweetsCount := work.MyTweetsCount
	if myTweetsCount >= 5 {
		myTweetsCount = 5
	}

	repliesCount := work.RepliesCount
	if repliesCount >= 10 {
		repliesCount = 10
	}

	favesCount := work.IncreaseFavoritesCount
	if favesCount >= 500 {
		favesCount = 500
	}

	tweetRate := (1.0 / 25.0) * math.Pow(float64(myTweetsCount), 2)
	replyRate := (1.0 / 100.0) * math.Pow(float64(repliesCount), 2)
	favesRate := (1.0 / 250000.0) * math.Pow(float64(favesCount), 2)

	totalPoint := 40*tweetRate + 30*replyRate + 30*favesRate

	work.Point = math.Round(totalPoint*100) / 100
}
