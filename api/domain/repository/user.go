//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=../../mock/$GOPACKAGE/mock_$GOFILE

package repository

import "influ-dojo/api/domain/model"

type User interface {
	LoadByID(userID string) (*model.User, error)
	LoadByScreenName(screenName string) (*model.User, error)
	Save(user *model.User) error
}
