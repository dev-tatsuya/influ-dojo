//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=../../mock/$GOPACKAGE/mock_$GOFILE

package client

import "influ-dojo/api/usecase/dto"

type Bot interface {
	Tweet(top3 *dto.Top3, path string) error
	Favorite() error
}
