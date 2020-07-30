package model

type DailyResult struct {
	ID                     int `gorm:"primary_key;auto_increment"`
	ScreenName             string
	FollowersCount         int
	IncreaseFollowersCount *int
	Point                  *int
	Model
}

func (mdl *DailyResult) IsNew() bool {
	return mdl.ID == 0
}

func (mdl *DailyResult) AttachID() error {
	return nil
}
