package client

import (
	"fmt"
	"influ-dojo/api/domain/client"
	"influ-dojo/api/domain/model"
	"log"
	"math"
	"math/rand"
	"net/url"
	"path"
	"strconv"
	"time"
)

type bot struct {
	*twitter
}

var period = map[string]string{
	"daily":   "デイリー",
	"weekly":  "ウィークリー",
	"monthly": "マンスリー",
}

func NewBot(accessToken, accessTokenSecret, consumerKey, consumerSecret string) client.Bot {
	return &bot{
		twitter: newTwitter(accessToken, accessTokenSecret, consumerKey, consumerSecret),
	}
}

func (b *bot) Tweet(works []*model.Work, results []*model.Result, pathStr string) error {
	body := fmt.Sprintf(`
【%s速報 %v】

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
		period[path.Base(pathStr)], time.Now().Format("1/2 15:04"),
		works[0].Point, works[0].ScreenName, works[1].Point, works[1].ScreenName, works[2].Point, works[2].ScreenName,
		results[0].Point, results[0].ScreenName, results[1].Point, results[1].ScreenName, results[2].Point, results[2].ScreenName,
	)

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
