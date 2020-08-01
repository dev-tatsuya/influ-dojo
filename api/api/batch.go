package api

import (
	"influ-dojo/api/interface/http/handler"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func NewCron(dependency *Dependency) {
	c := cron.New(cron.WithLocation(time.Local))

	setScheduler(c, dependency)

	c.Start()
}

func setScheduler(c *cron.Cron, d *Dependency) {
	setBatchSequence(c, "0 19 * * *", func() {
		log.Println("daily ranking batch")
		handler.MakeRecordHandler(d.FollowerClient, d.UserRepo, d.DailyWorkRepo, d.DailyResultRepo)
		handler.MakeRankingHandler(d.DailyWorkRepo, d.DailyResultRepo)
		handler.MakeCacheHandler(d.RankingQuery, d.RankingRepo)
		handler.MakeTweetHandler(d.BotClient, d.DailyWorkRepo, d.DailyResultRepo)
	})

	setBatchSequence(c, "0 20 * * 0", func() {
		log.Println("weekly ranking batch")
		handler.MakeRecordHandler(d.FollowerClient, d.UserRepo, d.WeeklyWorkRepo, d.WeeklyResultRepo)
		handler.MakeRankingHandler(d.WeeklyWorkRepo, d.WeeklyResultRepo)
		handler.MakeCacheHandler(d.RankingQuery, d.RankingRepo)
		handler.MakeTweetHandler(d.BotClient, d.WeeklyWorkRepo, d.WeeklyResultRepo)
	})

	setBatchSequence(c, "0 21 1 * *", func() {
		log.Println("monthly ranking batch")
		handler.MakeRecordHandler(d.FollowerClient, d.UserRepo, d.MonthlyWorkRepo, d.MonthlyResultRepo)
		handler.MakeRankingHandler(d.MonthlyWorkRepo, d.MonthlyResultRepo)
		handler.MakeCacheHandler(d.RankingQuery, d.RankingRepo)
		handler.MakeTweetHandler(d.BotClient, d.MonthlyWorkRepo, d.MonthlyResultRepo)
	})
}

func setBatchSequence(c *cron.Cron, spec string, f func()) {
	if _, err := c.AddFunc(spec, func() {
		f()
	}); err != nil {
		log.Printf("cron error: %+v", err)
	}
}
