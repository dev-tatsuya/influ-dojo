package api

import (
	domainClient "influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
	infraClient "influ-dojo/api/infrastructure/client"
	persistence "influ-dojo/api/infrastructure/persistence/repository"

	"github.com/jinzhu/gorm"
)

type Dependency struct {
	FollowerClient domainClient.Follower
	UserRepo       repository.User
	DailyWorkRepo  repository.DailyWork
}

func Inject(cfg *Config, db *gorm.DB) (*Dependency, error) {
	return &Dependency{
		FollowerClient: infraClient.NewFollower(
			cfg.Twitter.AccessToken,
			cfg.Twitter.AccessTokenSecret,
			cfg.Twitter.ConsumerKey,
			cfg.Twitter.ConsumerSecret,
		),
		UserRepo:      persistence.NewUser(db),
		DailyWorkRepo: persistence.NewDailyWork(db),
	}, nil
}
