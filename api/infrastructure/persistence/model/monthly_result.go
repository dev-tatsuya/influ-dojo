package model

type MonthlyResult struct {
	ID                     int `gorm:"primary_key;auto_increment"`
	ScreenName             string
	FollowersCount         int
	IncreaseFollowersCount *int
	Point                  *int
	Model
}

func (mdl *MonthlyResult) IsNew() bool {
	return mdl.ID == 0
}

func (mdl *MonthlyResult) AttachID() error {
	return nil
}
