package api

import (
	domainClient "influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
	infraClient "influ-dojo/api/infrastructure/client"
	persistence "influ-dojo/api/infrastructure/persistence/repository"

	"github.com/jinzhu/gorm"
)

type Dependency struct {
	FollowerClient    domainClient.Follower
	BotClient         domainClient.Bot
	UserRepo          repository.User
	DailyWorkRepo     repository.DailyWork
	DailyResultRepo   repository.DailyResult
	WeeklyWorkRepo    repository.WeeklyWork
	WeeklyResultRepo  repository.WeeklyResult
	MonthlyWorkRepo   repository.MonthlyWork
	MonthlyResultRepo repository.MonthlyResult
}

func Inject(cfg *Config, db *gorm.DB) (*Dependency, error) {
	at := cfg.Twitter.AccessToken
	ats := cfg.Twitter.AccessTokenSecret
	ck := cfg.Twitter.ConsumerKey
	cs := cfg.Twitter.ConsumerSecret

	return &Dependency{
		FollowerClient:    infraClient.NewFollower(at, ats, ck, cs),
		BotClient:         infraClient.NewBot(at, ats, ck, cs),
		UserRepo:          persistence.NewUser(db),
		DailyWorkRepo:     persistence.NewDailyWork(db),
		DailyResultRepo:   persistence.NewDailyResult(db),
		WeeklyWorkRepo:    persistence.NewWeeklyWork(db),
		WeeklyResultRepo:  persistence.NewDailyResult(db),
		MonthlyWorkRepo:   persistence.NewMonthlyWork(db),
		MonthlyResultRepo: persistence.NewMonthlyResult(db),
	}, nil
}
