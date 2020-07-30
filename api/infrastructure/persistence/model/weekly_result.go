package model

type WeeklyResult struct {
	ID                     int `gorm:"primary_key;auto_increment"`
	ScreenName             string
	FollowersCount         int
	IncreaseFollowersCount *int
	Point                  *int
	Model
}

func (mdl *WeeklyResult) IsNew() bool {
	return mdl.ID == 0
}

func (mdl *WeeklyResult) AttachID() error {
	return nil
}
