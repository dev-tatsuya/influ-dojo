package model

type Work struct {
	ScreenName             string
	TweetsCount            int
	IncreaseTweetsCount    int
	FavoritesCount         int
	IncreaseFavoritesCount int
	Point                  int
}

func (work *Work) SetPoint() {
	work.Point = work.IncreaseTweetsCount*200 + work.IncreaseFavoritesCount
}
