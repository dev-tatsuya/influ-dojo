package input

import (
	"influ-dojo/api/domain/apperr"
	domainClient "influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/output"
)

type DailyRank struct {
	FollowerClient  domainClient.Follower  `json:"-"`
	UserRepo        repository.User        `json:"-"`
	DailyWorkRepo   repository.DailyWork   `json:"-"`
	DailyResultRepo repository.DailyResult `json:"-"`
}

func (dr *DailyRank) GetDailyRank() (*output.DailyRank, error) {
	followers, err := dr.FollowerClient.GetFollowers()
	if err != nil {
		return nil, err
	}

	workUsers := make([]*output.WorkUser, 0)
	resultUsers := make([]*output.ResultUser, 0)
	for _, f := range followers {
		if f == nil {
			continue
		}

		if _, err := dr.UserRepo.LoadByID(f.User.UserID); err != nil {
			if err == apperr.ErrRecordNotFound {
				if err := dr.UserRepo.Save(f.User); err != nil {
					return nil, err
				}

				if err := dr.DailyWorkRepo.Save(f.Work); err != nil {
					return nil, err
				}

				if err := dr.DailyResultRepo.Save(f.Result); err != nil {
					return nil, err
				}

				continue
			}

			return nil, err
		}

		work, err := dr.DailyWorkRepo.LoadByID(f.User.UserID)
		if err != nil {
			return nil, err
		}

		result, err := dr.DailyResultRepo.LoadByID(f.User.UserID)
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

		resultUser := &output.ResultUser{
			Name:                   f.Name,
			ScreenName:             f.ScreenName,
			ProfileImage:           f.ProfileImage,
			IncreaseFollowersCount: f.FollowersCount - result.FollowersCount,
		}

		workUsers = append(workUsers, workUser)
		resultUsers = append(resultUsers, resultUser)

		if err := dr.DailyWorkRepo.Save(f.Work); err != nil {
			return nil, err
		}

		if err := dr.DailyResultRepo.Save(f.Result); err != nil {
			return nil, err
		}
	}

	return &output.DailyRank{
		WorkRank:   &output.WorkRank{WorkUsers: workUsers},
		ResultRank: &output.ResultRank{ResultUsers: resultUsers},
	}, nil
}
