package model

type MonthlyResult struct {
	ScreenName             string `gorm:"primary_key"`
	FollowersCount         int
	IncreaseFollowersCount *int
	Point                  *int
	Model
}

func (mdl *MonthlyResult) IsNew() bool {
	return len(mdl.ScreenName) == 0
}

func (mdl *MonthlyResult) AttachID() error {
	return nil
}
