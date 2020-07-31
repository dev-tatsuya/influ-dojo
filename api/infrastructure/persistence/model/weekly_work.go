package model

import "influ-dojo/api/domain/model"

type WeeklyWork struct {
	ScreenName             string `gorm:"primary_key"`
	TweetsCount            int
	IncreaseTweetsCount    *int
	FavoritesCount         int
	IncreaseFavoritesCount *int
	Point                  *int
	Ranking                int
	LastRanking            int
	Model
}

func (mdl *WeeklyWork) IsNew() bool {
	return len(mdl.ScreenName) == 0
}

func (mdl *WeeklyWork) AttachID() error {
	return nil
}

func (mdl *WeeklyWork) MakeEntity() *model.Work {
	count := 0
	if mdl.IncreaseTweetsCount == nil {
		mdl.IncreaseTweetsCount = &count
	}
	if mdl.IncreaseFavoritesCount == nil {
		mdl.IncreaseFavoritesCount = &count
	}
	if mdl.Point == nil {
		mdl.Point = &count
	}

	return &model.Work{
		ScreenName:             mdl.ScreenName,
		TweetsCount:            mdl.TweetsCount,
		IncreaseTweetsCount:    *mdl.IncreaseTweetsCount,
		FavoritesCount:         mdl.FavoritesCount,
		IncreaseFavoritesCount: *mdl.IncreaseFavoritesCount,
		Point:                  *mdl.Point,
		Ranking:                mdl.Ranking,
		LastRanking:            mdl.LastRanking,
	}
}
