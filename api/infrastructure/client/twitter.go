package client

import (
	"github.com/ChimeraCoder/anaconda"
)

type twitter struct {
	api *anaconda.TwitterApi
}

func newTwitter(accessToken, accessTokenSecret, consumerKey, consumerSecret string) *twitter {
	return &twitter{
		api: anaconda.NewTwitterApiWithCredentials(
			accessToken, accessTokenSecret, consumerKey, consumerSecret,
		),
	}
}
