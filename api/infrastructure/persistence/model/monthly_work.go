package model

type MonthlyWork struct {
	ScreenName             string `gorm:"primary_key"`
	TweetsCount            int
	IncreaseTweetsCount    *int
	FavoritesCount         int
	IncreaseFavoritesCount *int
	Point                  *int
	Model
}

func (mdl *MonthlyWork) IsNew() bool {
	return len(mdl.ScreenName) == 0
}

func (mdl *MonthlyWork) AttachID() error {
	return nil
}
