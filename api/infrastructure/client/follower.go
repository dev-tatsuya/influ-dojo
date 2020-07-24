package client

import (
	"influ-dojo/api/domain/client"
	"net/url"
)

type follower struct {
	*twitter
}

func NewFollower(accessToken, accessTokenSecret, consumerKey, consumerSecret string) client.Follower {
	return &follower{
		newTwitter(accessToken, accessTokenSecret, consumerKey, consumerSecret),
	}
}

func (f follower) CountFollowers() (int, error) {
	values := url.Values{}
	values.Add("screen_name", "soldinx")
	cursor, err := f.api.GetFollowersIds(values)
	if err != nil {
		return 0, err
	}
	return len(cursor.Ids), nil
}

func (f follower) GetFollowers() ([]client.Follower, error) {
	return nil, nil
}
