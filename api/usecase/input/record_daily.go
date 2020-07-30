package input

import (
	"influ-dojo/api/domain/apperr"
	domainClient "influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
	"influ-dojo/api/domain/utils"
)

type DailyRecord struct {
	FollowerClient  domainClient.Follower `json:"-"`
	UserRepo        repository.User       `json:"-"`
	DailyWorkRepo   repository.Work       `json:"-"`
	DailyResultRepo repository.Result     `json:"-"`
}

func (in *DailyRecord) RecordDaily() error {
	followers, err := in.FollowerClient.GetFollowers()
	if err != nil {
		return err
	}

	for _, f := range followers {
		if f == nil {
			continue
		}

		user, err := in.UserRepo.LoadByID(f.User.UserID)
		if err != nil {
			if err == apperr.ErrRecordNotFound {
				if err := in.UserRepo.Save(f.User); err != nil {
					return err
				}

				if err := in.DailyWorkRepo.Save(f.Work); err != nil {
					return err
				}

				if err := in.DailyResultRepo.Save(f.Result); err != nil {
					return err
				}

				continue
			}

			return err
		}

		work, err := in.DailyWorkRepo.LoadByScreenName(user.ScreenName)
		if err != nil {
			if err == apperr.ErrRecordNotFound {
				if err := in.DailyWorkRepo.Save(f.Work); err != nil {
					return err
				}

				if err := in.DailyResultRepo.Save(f.Result); err != nil {
					return err
				}

				continue
			}

			return err
		}

		result, err := in.DailyResultRepo.LoadByScreenName(user.ScreenName)
		if err != nil {
			return err
		}

		f.Work.IncreaseTweetsCount = utils.Sub(f.TweetsCount, work.TweetsCount)
		f.Work.IncreaseFavoritesCount = utils.Sub(f.FavoritesCount, work.FavoritesCount)
		f.Work.SetPoint()

		f.Result.IncreaseFollowersCount = utils.Sub(f.FollowersCount, result.FollowersCount)
		f.Result.SetPoint()

		if err := in.DailyWorkRepo.Save(f.Work); err != nil {
			return err
		}

		if err := in.DailyResultRepo.Save(f.Result); err != nil {
			return err
		}
	}

	return nil
}
