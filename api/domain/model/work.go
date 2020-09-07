package model

import (
	"influ-dojo/api/domain/utils"
	"math"
)

type Work struct {
	UserID                 string
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

func (work *Work) CalcPoint(tweetBase, repBase, FavBase float64) {
	myTweetsCount := work.MyTweetsCount
	if myTweetsCount >= int(tweetBase) {
		myTweetsCount = int(tweetBase)
	} else if myTweetsCount <= 0 {
		myTweetsCount = 0
	}

	repliesCount := work.RepliesCount
	if repliesCount >= int(repBase) {
		repliesCount = int(repBase)
	} else if repliesCount <= 0 {
		repliesCount = 0
	}

	favesCount := work.IncreaseFavoritesCount
	if favesCount >= int(FavBase) {
		favesCount = int(FavBase)
	} else if favesCount <= 0 {
		favesCount = 0
	}

	tweetRate := (1.0 / (tweetBase * tweetBase)) * math.Pow(float64(myTweetsCount), 2)
	replyRate := (1.0 / (repBase * repBase)) * math.Pow(float64(repliesCount), 2)
	favesRate := (1.0 / (FavBase * FavBase)) * math.Pow(float64(favesCount), 2)

	totalPoint := 40*tweetRate + 30*replyRate + 30*favesRate

	work.Point = math.Round(totalPoint*100) / 100
}

func (work *Work) ResetTweetsCount() {
	work.MyTweetsCount = 0
	work.RepliesCount = 0
}
