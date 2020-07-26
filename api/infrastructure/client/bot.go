package client

import (
	"fmt"
	"influ-dojo/api/domain/client"
	"influ-dojo/api/domain/model"
	"log"
	"math"
	"math/rand"
	"net/url"
	"strconv"
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

func (b *bot) Tweet(works []*model.Work, results []*model.Result) error {
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
		works[3].Point, works[3].ScreenName, works[4].Point, works[4].ScreenName, works[5].Point, works[5].ScreenName,
		results[3].Point, results[3].ScreenName, results[4].Point, results[4].ScreenName, results[5].Point, results[5].ScreenName,
	)
	// TODO スライス3つ取得したのに、lengthが6になってるという謎

	if _, err := b.api.PostTweet(body, nil); err != nil {
		return err
	}

	return nil
}

var words = []string{"#今日の積み上げ", "#駆け出しエンジニア", "#駆け出しエンジニアと繋がりたい",
	"#プログラミング学習", "#プログラミング初心者", "#progate"}

func (b *bot) Favorite() error {
	likeIDs := make([]int64, 0)
	values := url.Values{}
	values.Set("result_type", "recent") //TODO 効いてないっぽい

	for _, word := range words {
		res, err := b.api.GetSearch(word, values)
		if err != nil {
			return err
		}

		for _, status := range res.Statuses {
			id, err := strconv.Atoi(status.IdStr)
			if err != nil {
				continue
			}

			likeIDs = append(likeIDs, int64(id))
		}
	}

	for _, id := range likeIDs {
		//人間が操作するくらいの時間間隔を空ける
		duration := int64(math.Ceil((rand.Float64()*2.0 + 2.0) * 1000)) //3s±1s
		time.Sleep(time.Duration(duration) * time.Millisecond)
		if _, err := b.api.Favorite(id); err != nil {
			if err.Error() == "Get https://api.twitter.com/1.1/favorites/create.json returned status 403, {\"errors\":[{\"code\":139,\"message\":\"You have already favorited this status.\"}]}" {
				continue
			}

			log.Printf("failed to favorite: %+v", err)
			return err
		}
	}

	return nil
}
