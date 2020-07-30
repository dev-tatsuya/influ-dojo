package model

type WeeklyWork struct {
	ID                     int `gorm:"primary_key;auto_increment"`
	ScreenName             string
	TweetsCount            int
	IncreaseTweetsCount    *int
	FavoritesCount         int
	IncreaseFavoritesCount *int
	Point                  *int
	Model
}

func (mdl *WeeklyWork) IsNew() bool {
	return mdl.ID == 0
}

func (mdl *WeeklyWork) AttachID() error {
	return nil
}
