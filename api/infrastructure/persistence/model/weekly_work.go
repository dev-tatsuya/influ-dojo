package model

import "influ-dojo/api/domain/model"

type WeeklyWork struct {
	ScreenName             string `gorm:"primary_key"`
	TweetsCount            int
	IncreaseTweetsCount    *int
	MyTweetsCount          *int
	RepliesCount           *int
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
	zero := 0
	if mdl.IncreaseTweetsCount == nil {
		mdl.IncreaseTweetsCount = &zero
	}
	if mdl.MyTweetsCount == nil {
		mdl.MyTweetsCount = &zero
	}
	if mdl.RepliesCount == nil {
		mdl.RepliesCount = &zero
	}
	if mdl.IncreaseFavoritesCount == nil {
		mdl.IncreaseFavoritesCount = &zero
	}
	if mdl.Point == nil {
		mdl.Point = &zero
	}

	return &model.Work{
		ScreenName:             mdl.ScreenName,
		TweetsCount:            mdl.TweetsCount,
		IncreaseTweetsCount:    *mdl.IncreaseTweetsCount,
		MyTweetsCount:          *mdl.MyTweetsCount,
		RepliesCount:           *mdl.RepliesCount,
		FavoritesCount:         mdl.FavoritesCount,
		IncreaseFavoritesCount: *mdl.IncreaseFavoritesCount,
		Point:                  *mdl.Point,
		Ranking:                mdl.Ranking,
		LastRanking:            mdl.LastRanking,
	}
}

func NewWeeklyWork(entity *model.Work) *WeeklyWork {
	return &WeeklyWork{
		ScreenName:             entity.ScreenName,
		TweetsCount:            entity.TweetsCount,
		IncreaseTweetsCount:    &entity.IncreaseTweetsCount,
		MyTweetsCount:          &entity.MyTweetsCount,
		RepliesCount:           &entity.RepliesCount,
		FavoritesCount:         entity.FavoritesCount,
		IncreaseFavoritesCount: &entity.IncreaseFavoritesCount,
		Point:                  &entity.Point,
		Ranking:                entity.Ranking,
		LastRanking:            entity.LastRanking,
	}
}
