package client

import (
	"fmt"
	"influ-dojo/api/domain/client"
	"influ-dojo/api/domain/model"
	"time"
)

type bot struct {
	*twitter
}

func NewBot(accessToken, accessTokenSecret, consumerKey, consumerSecret string) client.Bot {
	return &bot{
		newTwitter(accessToken, accessTokenSecret, consumerKey, consumerSecret),
	}
}

func (b bot) Tweet(works []*model.Work, results []*model.Result) error {
	fmt.Println("===== work ranking =====", len(works))
	for _, work := range works {
		if work == nil {
			continue
		}
		fmt.Println(*work)
	}

	fmt.Println("===== result ranking =====", len(results))
	for _, result := range results {
		if result == nil {
			continue
		}
		fmt.Println(*result)
	}

	body := fmt.Sprintf(`
【デイリー速報 %v】

［作業ランキング］
🥇 %dpt @%s
🥈 %dpt @%s
🥉 %dpt @%s

［成果ランキング］
🥇 %dpt @%s
🥈 %dpt @%s
🥉 %dpt @%s

▼ 4位以下はコチラ ▼
example.com
`,
		time.Now().Format("1/2 15:04:05"),
		works[3].Point, works[3].UserID, works[4].Point, works[4].UserID, works[5].Point, works[5].UserID,
		results[3].Point, results[3].UserID, results[4].Point, results[4].UserID, results[5].Point, results[5].UserID,
	)
	// TODO スライス3つ取得したのに、lengthが6になってるという謎

	if _, err := b.api.PostTweet(body, nil); err != nil {
		return err
	}

	return nil
}
