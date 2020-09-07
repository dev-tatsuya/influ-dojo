package client

import (
	"fmt"
	"influ-dojo/api/domain/client"
	"influ-dojo/api/domain/model"
	"influ-dojo/api/domain/utils"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
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
	values.Add("screen_name", "infludojo")
	cursor, err := f.api.GetFollowersIds(values)
	if err != nil {
		return 0, err
	}

	return len(cursor.Ids), nil
}

func (f follower) GetFollowers() ([]*model.Follower, error) {
	values := url.Values{}
	values.Add("screen_name", "infludojo")
	cursor, err := f.api.GetFollowersIds(values)
	if err != nil {
		return nil, err
	}

	totalUsers := make([]anaconda.User, 0)
	for index := range utils.IndexChunks(len(cursor.Ids), 100) {
		users, err := f.api.GetUsersLookupByIds(cursor.Ids[index.From:index.To], nil)
		if err != nil {
			return nil, err
		}

		totalUsers = append(totalUsers, users...)
	}

	fmt.Println("followers count:", len(totalUsers))

	followers := make([]*model.Follower, 0)
	for _, f := range totalUsers {
		f := &model.Follower{
			User: &model.User{
				UserID:       f.IdStr,
				Name:         f.Name,
				ScreenName:   f.ScreenName,
				ProfileImage: f.ProfileImageUrlHttps,
			},
			Work: &model.Work{
				UserID:         f.IdStr,
				TweetsCount:    int(f.StatusesCount),
				FavoritesCount: f.FavouritesCount,
			},
			Result: &model.Result{
				UserID:         f.IdStr,
				FollowersCount: f.FollowersCount,
			},
		}
		followers = append(followers, f)
	}

	return followers, nil
}
