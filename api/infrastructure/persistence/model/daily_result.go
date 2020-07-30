package model

import "influ-dojo/api/domain/model"

type DailyResult struct {
	ID                     int `gorm:"primary_key;auto_increment"`
	ScreenName             string
	FollowersCount         int
	IncreaseFollowersCount *int
	Point                  *int
	Ranking                int
	LastRanking            int
	Model
}

func (mdl *DailyResult) IsNew() bool {
	return mdl.ID == 0
}

func (mdl *DailyResult) AttachID() error {
	return nil
}

func (mdl *DailyResult) MakeEntity() *model.Result {
	count := 0
	if mdl.IncreaseFollowersCount == nil {
		mdl.IncreaseFollowersCount = &count
	}
	if mdl.Point == nil {
		mdl.Point = &count
	}

	return &model.Result{
		ScreenName:             mdl.ScreenName,
		FollowersCount:         mdl.FollowersCount,
		IncreaseFollowersCount: *mdl.IncreaseFollowersCount,
		Point:                  *mdl.Point,
		Ranking:                mdl.Ranking,
		LastRanking:            mdl.LastRanking,
	}
}
