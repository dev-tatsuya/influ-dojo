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
		handler.Record(d.FollowerClient, d.UserRepo, d.DailyWorkRepo, d.DailyResultRepo)
		handler.Ranking(d.DailyWorkRepo, d.DailyResultRepo)
		handler.Cache(d.RankingQuery, d.RankingRepo)
		handler.Tweet(d.BotClient, d.DailyWorkRepo, d.DailyResultRepo, "api/tweet/daily")
	})

	setBatchSequence(c, "0 20 * * 0", func() {
		log.Println("weekly ranking batch")
		handler.Record(d.FollowerClient, d.UserRepo, d.WeeklyWorkRepo, d.WeeklyResultRepo)
		handler.Ranking(d.WeeklyWorkRepo, d.WeeklyResultRepo)
		handler.Cache(d.RankingQuery, d.RankingRepo)
		handler.Tweet(d.BotClient, d.WeeklyWorkRepo, d.WeeklyResultRepo, "api/tweet/weekly")
	})

	setBatchSequence(c, "0 21 1 * *", func() {
		log.Println("monthly ranking batch")
		handler.Record(d.FollowerClient, d.UserRepo, d.MonthlyWorkRepo, d.MonthlyResultRepo)
		handler.Ranking(d.MonthlyWorkRepo, d.MonthlyResultRepo)
		handler.Cache(d.RankingQuery, d.RankingRepo)
		handler.Tweet(d.BotClient, d.MonthlyWorkRepo, d.MonthlyResultRepo, "api/tweet/monthly")
	})

	//setBatchSequence(c, "0 7 * * *", func() { handler.Favorite(d.BotClient) })
	//setBatchSequence(c, "0 12 * * *", func() { handler.Favorite(d.BotClient) })
	//setBatchSequence(c, "0 17 * * *", func() { handler.Favorite(d.BotClient) })
	//setBatchSequence(c, "0 18 * * *", func() { handler.Favorite(d.BotClient) })
	//setBatchSequence(c, "0 19 * * *", func() { handler.Favorite(d.BotClient) })
	//setBatchSequence(c, "0 20 * * *", func() { handler.Favorite(d.BotClient) })
	//setBatchSequence(c, "0 21 * * *", func() { handler.Favorite(d.BotClient) })
	//setBatchSequence(c, "0 22 * * *", func() { handler.Favorite(d.BotClient) })
	//setBatchSequence(c, "0 23 * * *", func() { handler.Favorite(d.BotClient) })
	//setBatchSequence(c, "0 24 * * *", func() { handler.Favorite(d.BotClient) })
}

func setBatchSequence(c *cron.Cron, spec string, f func()) {
	if _, err := c.AddFunc(spec, func() {
		f()
	}); err != nil {
		log.Printf("cron error: %+v", err)
	}
}
