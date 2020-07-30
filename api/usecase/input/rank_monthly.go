package input

import (
	"influ-dojo/api/domain/apperr"
	domainClient "influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/domain/utils"
)

type MonthlyRank struct {
	FollowerClient    domainClient.Follower `json:"-"`
	UserRepo          repository.User       `json:"-"`
	MonthlyWorkRepo   repository.Work       `json:"-"`
	MonthlyResultRepo repository.Result     `json:"-"`
}

func (rank *MonthlyRank) PostMonthlyRank() error {
	followers, err := rank.FollowerClient.GetFollowers()
	if err != nil {
		return err
	}

	for _, f := range followers {
		if f == nil {
			continue
		}

		user, err := rank.UserRepo.LoadByID(f.User.UserID)
		if err != nil {
			if err == apperr.ErrRecordNotFound {
				if err := rank.UserRepo.Save(f.User); err != nil {
					return err
				}

				if err := rank.MonthlyWorkRepo.Save(f.Work); err != nil {
					return err
				}

				if err := rank.MonthlyResultRepo.Save(f.Result); err != nil {
					return err
				}

				continue
			}

			return err
		}

		work, err := rank.MonthlyWorkRepo.LoadByScreenName(user.ScreenName)
		if err != nil {
			if err == apperr.ErrRecordNotFound {
				if err := rank.MonthlyWorkRepo.Save(f.Work); err != nil {
					return err
				}

				if err := rank.MonthlyResultRepo.Save(f.Result); err != nil {
					return err
				}

				continue
			}

			return err
		}

		result, err := rank.MonthlyResultRepo.LoadByScreenName(user.ScreenName)
		if err != nil {
			return err
		}

		f.Work.IncreaseTweetsCount = utils.Sub(f.TweetsCount, work.TweetsCount)
		f.Work.IncreaseFavoritesCount = utils.Sub(f.FavoritesCount, work.FavoritesCount)
		f.Work.SetPoint()

		f.Result.IncreaseFollowersCount = utils.Sub(f.FollowersCount, result.FollowersCount)
		f.Result.SetPoint()

		if err := rank.MonthlyWorkRepo.Save(f.Work); err != nil {
			return err
		}

		if err := rank.MonthlyResultRepo.Save(f.Result); err != nil {
			return err
		}
	}

	return nil
}
