package model

type WeeklyResult struct {
	ScreenName             string `gorm:"primary_key"`
	FollowersCount         int
	IncreaseFollowersCount *int
	Point                  *int
	Model
}

func (mdl *WeeklyResult) IsNew() bool {
	return len(mdl.ScreenName) == 0
}

func (mdl *WeeklyResult) AttachID() error {
	return nil
}
