package model

type User struct {
	UserID       string `gorm:"primary_key"`
	Name         string
	ScreenName   string
	ProfileImage string
	Model
}

func (mdl *User) IsNew() bool {
	return len(mdl.UserID) == 0
}

func (mdl *User) AttachID() error {
	return nil
}
