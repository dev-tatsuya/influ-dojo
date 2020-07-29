package model

type Result struct {
	ScreenName             string
	FollowersCount         int
	IncreaseFollowersCount int
	Point                  int
}

func (res *Result) SetPoint() {
	res.Point = res.IncreaseFollowersCount
}
