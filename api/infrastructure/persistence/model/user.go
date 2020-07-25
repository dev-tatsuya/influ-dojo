package model

type User struct {
	UserID       string `gorm:"primary_key"`
	Name         string
	ScreenName   string `gorm:"primary_key"`
	ProfileImage string
	Model
}
