//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=../../mock/$GOPACKAGE/mock_$GOFILE

package repository

import "influ-dojo/api/domain/model"

type DailyResult interface {
	Load() ([]*model.Result, error)
	LoadByID(userID string) (*model.Result, error)
	Save(work *model.Result) error
}
