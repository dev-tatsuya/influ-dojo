package input

import (
	"influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
	"log"
)

type ClassifyTweets struct {
	TweetClient     client.Tweet
	UserRepo        repository.User
	DailyWorkRepo   repository.Work
	WeeklyWorkRepo  repository.Work
	MonthlyWorkRepo repository.Work
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

		user, err := in.UserRepo.LoadByID(work.UserID)
		if err != nil {
			return err
		}

		screenName := user.ScreenName
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

		if err := addTweetsCount(in.WeeklyWorkRepo, user.UserID, myTweetCount, repliesCount); err != nil {
			return err
		}

		if err := addTweetsCount(in.MonthlyWorkRepo, user.UserID, myTweetCount, repliesCount); err != nil {
			return err
		}
	}

	return nil
}

func addTweetsCount(repo repository.Work, userID string, myTweetCount, repliesCount int) error {
	entity, err := repo.LoadByID(userID)
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
