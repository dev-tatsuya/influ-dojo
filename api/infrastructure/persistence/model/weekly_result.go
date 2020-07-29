package model

type WeeklyResult struct {
	ID                     int `gorm:"primary_key;auto_increment"`
	ScreenName             string
	FollowersCount         int
	IncreaseFollowersCount *int
	Point                  *int
	Model
}
