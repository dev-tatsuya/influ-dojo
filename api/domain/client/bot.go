//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=../../mock/$GOPACKAGE/mock_$GOFILE

package client

import "influ-dojo/api/domain/model"

type Bot interface {
	Tweet(work []*model.Work, result []*model.Result, path string) error
	Favorite() error
}
