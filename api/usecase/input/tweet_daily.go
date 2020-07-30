package input

import (
	"influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
)

type DailyTweet struct {
	Bot             client.Bot        `json:"-"`
	DailyWorkRepo   repository.Work   `json:"-"`
	DailyResultRepo repository.Result `json:"-"`
}

func (dt *DailyTweet) TweetDaily() error {
	works, err := dt.DailyWorkRepo.LoadTop3()
	if err != nil {
		return err
	}

	results, err := dt.DailyResultRepo.LoadTop3()
	if err != nil {
		return err
	}

	return dt.Bot.Tweet(works, results)
}
