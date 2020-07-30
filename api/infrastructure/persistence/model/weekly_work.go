package model

type WeeklyWork struct {
	ScreenName             string `gorm:"primary_key"`
	TweetsCount            int
	IncreaseTweetsCount    *int
	FavoritesCount         int
	IncreaseFavoritesCount *int
	Point                  *int
	Model
}

func (mdl *WeeklyWork) IsNew() bool {
	return len(mdl.ScreenName) == 0
}

func (mdl *WeeklyWork) AttachID() error {
	return nil
}
