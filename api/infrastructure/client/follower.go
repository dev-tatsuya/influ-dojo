package client

import (
	"fmt"
	"influ-dojo/api/domain/client"
	"influ-dojo/api/domain/model"
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

func (f follower) GetFollowers() ([]*model.Follower, error) {
	values := url.Values{}
	values.Add("screen_name", "soldinx")
	cursor, err := f.api.GetFriendsIds(values) //TODO フォロワーを取得すると多すぎる。クエリが膨大になる問題
	if err != nil {
		return nil, err
	}

	users, err := f.api.GetUsersLookupByIds(cursor.Ids, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("followers count:", len(users))

	followers := make([]*model.Follower, len(users))
	for _, f := range users {
		f := &model.Follower{
			User: &model.User{
				UserID:       f.IdStr,
				Name:         f.Name,
				ScreenName:   f.ScreenName,
				ProfileImage: f.ProfileImageURL,
			},
			Work: &model.Work{
				ScreenName:     f.ScreenName,
				TweetsCount:    int(f.StatusesCount),
				FavoritesCount: f.FavouritesCount,
			},
			Result: &model.Result{
				ScreenName:     f.ScreenName,
				FollowersCount: f.FollowersCount,
			},
		}
		followers = append(followers, f)
	}

	return followers, nil
}
