package input

import (
	domainClient "influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
	"log"
)

type Record struct {
	FollowerClient domainClient.Follower `json:"-"`
	UserRepo       repository.User       `json:"-"`
	WorkRepo       repository.Work       `json:"-"`
	ResultRepo     repository.Result     `json:"-"`
}

func (in *Record) Record() error {
	followers, err := in.FollowerClient.GetFollowers()
	if err != nil {
		return err
	}

	for _, f := range followers {
		//TODO フォロワーの数だけDB接続してしまう。+QueryServiceでまとめて取得したい
		user, err := in.UserRepo.LoadByID(f.User.UserID)
		if err != nil {
			log.Printf("failed to load user by %s: error: %+v", f.UserID, err)
			continue
		}

		work, err := in.WorkRepo.LoadByScreenName(user.ScreenName)
		if err != nil {
			log.Printf("failed to load wok by %s: error: %+v", user.ScreenName, err)
			continue
		}

		result, err := in.ResultRepo.LoadByScreenName(user.ScreenName)
		if err != nil {
			log.Printf("failed to load result by %s: error: %+v", user.ScreenName, err)
			continue
		}

		work.UpdateCount(f.TweetsCount, f.FavoritesCount)

		result.UpdateCount(f.FollowersCount)

		if user.IsUpdateRequired(f.Name, f.User.ScreenName, f.ProfileImage) {
			if err := in.UserRepo.Save(user); err != nil {
				return err
			}
		}

		if err := in.WorkRepo.Save(work); err != nil {
			return err
		}

		if err := in.ResultRepo.Save(result); err != nil {
			return err
		}
	}

	return nil
}
