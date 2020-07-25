package model

type DailyWork struct {
	ID             int `gorm:"primary_key;auto_increment"`
	UserID         string
	TweetsCount    int
	FavoritesCount int
	Model
}
