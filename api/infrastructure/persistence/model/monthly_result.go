package model

import "influ-dojo/api/domain/model"

type MonthlyResult struct {
	UserID                 string `gorm:"primary_key"`
	FollowersCount         int
	IncreaseFollowersCount *int
	Point                  *int
	Ranking                int
	LastRanking            int
	Model
}

func (mdl *MonthlyResult) IsNew() bool {
	return len(mdl.UserID) == 0
}

func (mdl *MonthlyResult) AttachID() error {
	return nil
}

func (mdl *MonthlyResult) MakeEntity() *model.Result {
	count := 0
	if mdl.IncreaseFollowersCount == nil {
		mdl.IncreaseFollowersCount = &count
	}
	if mdl.Point == nil {
		mdl.Point = &count
	}

	return &model.Result{
		UserID:                 mdl.UserID,
		FollowersCount:         mdl.FollowersCount,
		IncreaseFollowersCount: *mdl.IncreaseFollowersCount,
		Point:                  *mdl.Point,
		Ranking:                mdl.Ranking,
		LastRanking:            mdl.LastRanking,
	}
}
