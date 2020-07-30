package model

type MonthlyWork struct {
	ID                     int `gorm:"primary_key;auto_increment"`
	ScreenName             string
	TweetsCount            int
	IncreaseTweetsCount    *int
	FavoritesCount         int
	IncreaseFavoritesCount *int
	Point                  *int
	Model
}

func (mdl *MonthlyWork) IsNew() bool {
	return mdl.ID == 0
}

func (mdl *MonthlyWork) AttachID() error {
	return nil
}
