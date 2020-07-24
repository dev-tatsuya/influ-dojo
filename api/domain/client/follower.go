//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=../../mock/$GOPACKAGE/mock_$GOFILE

package client

type Follower interface {
	CountFollowers() (int, error)
	GetFollowers() ([]Follower, error)
}
