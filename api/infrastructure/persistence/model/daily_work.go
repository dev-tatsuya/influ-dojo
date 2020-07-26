package model

type DailyWork struct {
	ID                     int `gorm:"primary_key;auto_increment"`
	ScreenName             string
	TweetsCount            int
	IncreaseTweetsCount    *int
	FavoritesCount         int
	IncreaseFavoritesCount *int
	Point                  *int
	Model
}
