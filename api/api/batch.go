package api

import (
	"fmt"
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
	setBatchSequence(c, "@every 30m", func() {
		fmt.Println("register new users")
		handler.RecordNewUsers(d.FollowerClient, d.UserRepo, d.DailyWorkRepo, d.DailyResultRepo,
			d.WeeklyWorkRepo, d.WeeklyResultRepo, d.MonthlyWorkRepo, d.MonthlyResultRepo)
	})

	setBatchSequence(c, "0 19 * * *", func() {
		fmt.Println("daily ranking batch")
		handler.RecordNewUsers(d.FollowerClient, d.UserRepo, d.DailyWorkRepo, d.DailyResultRepo,
			d.WeeklyWorkRepo, d.WeeklyResultRepo, d.MonthlyWorkRepo, d.MonthlyResultRepo)
		handler.Record(d.FollowerClient, d.UserRepo, d.DailyWorkRepo, d.DailyResultRepo)
		handler.ClassifyDailyTweets(d.TweetClient, d.DailyWorkRepo, d.WeeklyWorkRepo, d.MonthlyWorkRepo)
		handler.CalcPoint(d.DailyWorkRepo)
		handler.Ranking(d.DailyWorkRepo, d.DailyResultRepo)
		handler.Cache(d.RankingQuery, d.RankingRepo)
		handler.Tweet(d.BotClient, d.DailyWorkRepo, d.DailyResultRepo, "api/tweet/daily")
	})

	setBatchSequence(c, "0 20 * * 0", func() {
		fmt.Println("weekly ranking batch")
		handler.RecordNewUsers(d.FollowerClient, d.UserRepo, d.DailyWorkRepo, d.DailyResultRepo,
			d.WeeklyWorkRepo, d.WeeklyResultRepo, d.MonthlyWorkRepo, d.MonthlyResultRepo)
		handler.Record(d.FollowerClient, d.UserRepo, d.WeeklyWorkRepo, d.WeeklyResultRepo)
		handler.CalcPoint(d.WeeklyWorkRepo)
		handler.Ranking(d.WeeklyWorkRepo, d.WeeklyResultRepo)
		handler.Cache(d.RankingQuery, d.RankingRepo)
		handler.Tweet(d.BotClient, d.WeeklyWorkRepo, d.WeeklyResultRepo, "api/tweet/weekly")
		handler.ResetTweetsCount(d.WeeklyWorkRepo)
	})

	setBatchSequence(c, "0 21 1 * *", func() {
		fmt.Println("monthly ranking batch")
		handler.RecordNewUsers(d.FollowerClient, d.UserRepo, d.DailyWorkRepo, d.DailyResultRepo,
			d.WeeklyWorkRepo, d.WeeklyResultRepo, d.MonthlyWorkRepo, d.MonthlyResultRepo)
		handler.Record(d.FollowerClient, d.UserRepo, d.MonthlyWorkRepo, d.MonthlyResultRepo)
		handler.CalcPoint(d.MonthlyWorkRepo)
		handler.Ranking(d.MonthlyWorkRepo, d.MonthlyResultRepo)
		handler.Cache(d.RankingQuery, d.RankingRepo)
		handler.Tweet(d.BotClient, d.MonthlyWorkRepo, d.MonthlyResultRepo, "api/tweet/monthly")
		handler.ResetTweetsCount(d.MonthlyWorkRepo)
	})

	setBatchSequence(c, "0 7 * * *", func() { handler.Favorite(d.BotClient) })
	setBatchSequence(c, "0 12 * * *", func() { handler.Favorite(d.BotClient) })
	setBatchSequence(c, "0 17 * * *", func() { handler.Favorite(d.BotClient) })
	setBatchSequence(c, "0 18 * * *", func() { handler.Favorite(d.BotClient) })
	setBatchSequence(c, "0 19 * * *", func() { handler.Favorite(d.BotClient) })
	setBatchSequence(c, "0 20 * * *", func() { handler.Favorite(d.BotClient) })
	setBatchSequence(c, "0 21 * * *", func() { handler.Favorite(d.BotClient) })
	setBatchSequence(c, "0 22 * * *", func() { handler.Favorite(d.BotClient) })
	setBatchSequence(c, "0 23 * * *", func() { handler.Favorite(d.BotClient) })
	setBatchSequence(c, "0 0 * * *", func() { handler.Favorite(d.BotClient) })
}

func setBatchSequence(c *cron.Cron, spec string, f func()) {
	if _, err := c.AddFunc(spec, func() {
		f()
	}); err != nil {
		log.Printf("cron error: %+v", err)
	}
}
