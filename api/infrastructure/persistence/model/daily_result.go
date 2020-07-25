package model

type DailyResult struct {
	ID             int `gorm:"primary_key;auto_increment"`
	UserID         string
	FollowersCount int
	Model
}
