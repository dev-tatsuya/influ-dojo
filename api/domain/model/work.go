package model

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

func (work *Work) SetPoint() {
	work.Point = work.IncreaseTweetsCount*200 + work.IncreaseFavoritesCount
}

func (work *Work) MakeRankingPast() {
	work.LastRanking = work.Ranking
}
