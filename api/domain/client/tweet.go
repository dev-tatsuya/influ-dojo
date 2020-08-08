//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=../../mock/$GOPACKAGE/mock_$GOFILE

package client

import "influ-dojo/api/domain/model"

type Tweet interface {
	FetchTweetsFromScreenName(screenName string, count int) ([]*model.Tweet, error)
}
