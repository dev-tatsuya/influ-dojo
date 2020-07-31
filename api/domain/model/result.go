package model

import "influ-dojo/api/domain/utils"

type Result struct {
	ScreenName             string
	FollowersCount         int
	IncreaseFollowersCount int
	Point                  int
	Ranking                int
	LastRanking            int
}

func (res *Result) MakeRankingPast() {
	res.LastRanking = res.Ranking
}

func (res *Result) UpdateCount(latestFollowersCount int) {
	res.IncreaseFollowersCount = utils.Sub(latestFollowersCount, res.FollowersCount)
	res.setPoint()
	res.FollowersCount = latestFollowersCount
}

func (res *Result) setPoint() {
	res.Point = res.IncreaseFollowersCount
}
