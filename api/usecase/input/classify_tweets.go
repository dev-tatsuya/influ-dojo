package input

import (
	"influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
	"log"
)

type ClassifyTweets struct {
	TweetClient     client.Tweet    `json:"-"`
	DailyWorkRepo   repository.Work `json:"-"`
	WeeklyWorkRepo  repository.Work `json:"-"`
	MonthlyWorkRepo repository.Work `json:"-"`
}

func (in *ClassifyTweets) Classify() error {
	works, err := in.DailyWorkRepo.Load()
	if err != nil {
		return err
	}

	for _, work := range works {
		if work.IncreaseTweetsCount <= 0 {
			continue
		}

		screenName := work.ScreenName
		tweets, err := in.TweetClient.FetchTweetsFromScreenName(screenName, work.IncreaseTweetsCount)
		if err != nil {
			log.Printf("failed to fetch tweets from %s, error: %+v", screenName, err)
			continue
		}

		myTweetCount := 0
		repliesCount := 0
		for _, tweet := range tweets {
			if tweet.IsRetweetedStatus {
				continue
			}

			if len(tweet.InReplyToScreenName) != 0 && tweet.InReplyToScreenName != screenName {
				if len(tweet.Text) < 50 {
					continue
				}

				repliesCount++
				continue
			}

			if len(tweet.Text) < 100 {
				continue
			}

			myTweetCount++
		}

		work.MyTweetsCount = myTweetCount
		work.RepliesCount = repliesCount
		if err := in.DailyWorkRepo.Save(work); err != nil {
			return err
		}

		if err := addTweetsCount(in.WeeklyWorkRepo, screenName, myTweetCount, repliesCount); err != nil {
			return err
		}

		if err := addTweetsCount(in.MonthlyWorkRepo, screenName, myTweetCount, repliesCount); err != nil {
			return err
		}
	}

	return nil
}

func addTweetsCount(repo repository.Work, screenName string, myTweetCount, repliesCount int) error {
	entity, err := repo.LoadByScreenName(screenName)
	if err != nil {
		return err
	}

	if myTweetCount > 5 {
		entity.MyTweetsCount += 5
	} else {
		entity.MyTweetsCount += myTweetCount
	}

	if repliesCount > 10 {
		entity.RepliesCount += 10
	} else {
		entity.RepliesCount += repliesCount
	}

	return repo.Save(entity)
}
