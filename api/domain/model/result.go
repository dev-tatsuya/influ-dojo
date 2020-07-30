package model

type Result struct {
	ScreenName             string
	FollowersCount         int
	IncreaseFollowersCount int
	Point                  int
	Ranking                int
	LastRanking            int
}

func (res *Result) SetPoint() {
	res.Point = res.IncreaseFollowersCount
}

func (res *Result) MakeRankingPast() {
	res.LastRanking = res.Ranking
}
