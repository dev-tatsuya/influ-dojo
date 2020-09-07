//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=../../mock/$GOPACKAGE/mock_$GOFILE

package repository

import "influ-dojo/api/domain/model"

type Result interface {
	LoadOrderByRanking() ([]*model.Result, error)
	LoadByID(id string) (*model.Result, error)
	Save(work *model.Result) error
}
