//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=../../mock/$GOPACKAGE/mock_$GOFILE

package repository

import "influ-dojo/api/domain/model"

type DailyWork interface {
	LoadTop3() ([]*model.Work, error)
	LoadByScreenName(screenName string) (*model.Work, error)
	Save(work *model.Work) error
}
