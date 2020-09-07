package client

import (
	"fmt"
	"influ-dojo/api/domain/client"
	"influ-dojo/api/domain/model"
	"net/url"
	"strconv"
)

type tweet struct {
	*twitter
}

func NewTweet(accessToken, accessTokenSecret, consumerKey, consumerSecret string) client.Tweet {
	return &tweet{
		newTwitter(accessToken, accessTokenSecret, consumerKey, consumerSecret),
	}
}

func (client *tweet) FetchTweetsFromScreenName(screenName string, count int) ([]*model.Tweet, error) {
	if count > 100 {
		count = 100
	}

	values := url.Values{}
	values.Add("count", strconv.Itoa(count))
	//TODO 1日100ツイート以上を考慮していない
	sr, err := client.api.GetSearch(fmt.Sprintf("from:@%s", screenName), values)
	if err != nil {
		return nil, err
	}

	tweets := make([]*model.Tweet, 0)
	for _, status := range sr.Statuses {
		tweet := &model.Tweet{
			InReplyToScreenName: status.InReplyToScreenName,
			IsRetweetedStatus:   status.RetweetedStatus != nil,
			Text:                status.Text,
		}

		tweets = append(tweets, tweet)
	}

	return tweets, nil
}
