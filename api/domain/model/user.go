package model

type User struct {
	UserID       string
	Name         string
	ScreenName   string
	ProfileImage string
}

func (user *User) IsUpdateRequired(name, screenName, image string) bool {
	isUpdateRequired := false

	if user.Name != name {
		user.Name = name
		isUpdateRequired = true
	}
	if user.ScreenName != screenName {
		user.ScreenName = screenName
		isUpdateRequired = true
	}
	if user.ProfileImage != image {
		user.ProfileImage = image
		isUpdateRequired = true
	}

	return isUpdateRequired
}
