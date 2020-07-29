package input

import (
	"influ-dojo/api/domain/apperr"
	domainClient "influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/usecase/output"
)

type MonthlyRank struct {
	FollowerClient    domainClient.Follower    `json:"-"`
	UserRepo          repository.User          `json:"-"`
	MonthlyWorkRepo   repository.MonthlyWork   `json:"-"`
	MonthlyResultRepo repository.MonthlyResult `json:"-"`
}

func (dr *MonthlyRank) GetMonthlyRank() (*output.DailyRank, error) {
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

				if err := dr.MonthlyWorkRepo.Save(f.Work); err != nil {
					return nil, err
				}

				if err := dr.MonthlyResultRepo.Save(f.Result); err != nil {
					return nil, err
				}

				continue
			}

			return nil, err
		}

		work, err := dr.MonthlyWorkRepo.LoadByScreenName(f.User.ScreenName)
		if err != nil {
			return nil, err
		}

		result, err := dr.MonthlyResultRepo.LoadByScreenName(f.User.ScreenName)
		if err != nil {
			return nil, err
		}

		increaseTweetsCount := f.TweetsCount - work.TweetsCount
		increaseFavoritesCount := f.FavoritesCount - work.FavoritesCount
		increaseFollowersCount := f.FollowersCount - result.FollowersCount

		workPoint := increaseTweetsCount*200 + increaseFavoritesCount
		resultPoint := increaseFollowersCount

		workUser := &output.WorkUser{
			Name:                   f.Name,
			ScreenName:             f.Work.ScreenName,
			ProfileImage:           f.ProfileImage,
			IncreaseTweetsCount:    increaseTweetsCount,
			IncreaseFavoritesCount: increaseFavoritesCount,
			Point:                  workPoint,
		}

		resultUser := &output.ResultUser{
			Name:                   f.Name,
			ScreenName:             f.Result.ScreenName,
			ProfileImage:           f.ProfileImage,
			IncreaseFollowersCount: increaseFollowersCount,
			Point:                  resultPoint,
		}

		workUsers = append(workUsers, workUser)
		resultUsers = append(resultUsers, resultUser)

		f.Work.IncreaseTweetsCount = increaseTweetsCount
		f.Work.IncreaseFavoritesCount = increaseFavoritesCount
		f.Work.Point = workPoint

		f.Result.IncreaseFollowersCount = increaseFollowersCount
		f.Result.Point = resultPoint

		if err := dr.MonthlyWorkRepo.Save(f.Work); err != nil {
			return nil, err
		}

		if err := dr.MonthlyResultRepo.Save(f.Result); err != nil {
			return nil, err
		}
	}

	return &output.DailyRank{
		WorkRank:   &output.WorkRank{WorkUsers: workUsers},
		ResultRank: &output.ResultRank{ResultUsers: resultUsers},
	}, nil
}
