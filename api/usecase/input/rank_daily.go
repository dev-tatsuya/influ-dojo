package input

import (
	"fmt"
	"influ-dojo/api/domain/apperr"
	domainClient "influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/output"
)

type DailyRank struct {
	FollowerClient domainClient.Follower `json:"-"`
	UserRepo       repository.User       `json:"-"`
	DailyWorkRepo  repository.DailyWork  `json:"-"`
}

func (dr *DailyRank) GetDailyRank() (*output.DailyRank, error) {
	// Twitterから全部取得
	followers, err := dr.FollowerClient.GetFollowers()
	if err != nil {
		return nil, err
	}

	workUsers := make([]*output.WorkUser, 0)
	for _, f := range followers {
		if f == nil {
			continue
		}
		fmt.Printf("get user: %+v", *f)
		// UserIDがDBにあるかどうか
		if _, err := dr.UserRepo.LoadByID(f.User.UserID); err != nil {
			if err == apperr.ErrRecordNotFound {
				// ユーザなどもろもろ登録
				if err := dr.UserRepo.Save(f.User); err != nil {
					return nil, err
				}

				if err := dr.DailyWorkRepo.Save(f.Work); err != nil {
					return nil, err
				}

				continue
			}

			return nil, err
		}

		// DBからロードして増加分を格納後、最新の値をセーブ
		work, err := dr.DailyWorkRepo.LoadByID(f.User.UserID)
		if err != nil {
			return nil, err
		}

		workUser := &output.WorkUser{
			Name:                   f.Name,
			ScreenName:             f.ScreenName,
			ProfileImage:           f.ProfileImage,
			IncreaseTweetsCount:    f.TweetsCount - work.TweetsCount,
			IncreaseFavoritesCount: f.FavoritesCount - work.FavoritesCount,
		}

		workUsers = append(workUsers, workUser)

		if err := dr.DailyWorkRepo.Save(f.Work); err != nil {
			return nil, err
		}
	}

	return &output.DailyRank{
		WorkRank:   &output.WorkRank{WorkUsers: workUsers},
		ResultRank: nil,
	}, nil
}
