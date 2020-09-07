package client

import (
	"fmt"
	"influ-dojo/api/domain/client"
	"influ-dojo/api/usecase/dto"
	"math"
	"math/rand"
	"net/url"
	"strconv"
	"time"

	"golang.org/x/xerrors"
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

func (b *bot) Tweet(top3 *dto.Top3, path string) error {
	workPoint1 := 0.0
	workPoint2 := 0.0
	workPoint3 := 0.0
	workName1 := ""
	workName2 := ""
	workName3 := ""
	resultPoint1 := 0
	resultPoint2 := 0
	resultPoint3 := 0
	resultName1 := ""
	resultName2 := ""
	resultName3 := ""

	switch path {
	case "daily":
		workPoint1 = top3.DailyWorkUsers[0].Point
		workPoint2 = top3.DailyWorkUsers[1].Point
		workPoint3 = top3.DailyWorkUsers[2].Point
		workName1 = top3.DailyWorkUsers[0].ScreenName
		workName2 = top3.DailyWorkUsers[1].ScreenName
		workName3 = top3.DailyWorkUsers[2].ScreenName
		resultPoint1 = int(top3.DailyResultUsers[0].Point)
		resultPoint2 = int(top3.DailyResultUsers[1].Point)
		resultPoint3 = int(top3.DailyResultUsers[2].Point)
		resultName1 = top3.DailyResultUsers[0].ScreenName
		resultName2 = top3.DailyResultUsers[1].ScreenName
		resultName3 = top3.DailyResultUsers[2].ScreenName
	case "weekly":
		workPoint1 = top3.WeeklyWorkUsers[0].Point
		workPoint2 = top3.WeeklyWorkUsers[1].Point
		workPoint3 = top3.WeeklyWorkUsers[2].Point
		workName1 = top3.WeeklyWorkUsers[0].ScreenName
		workName2 = top3.WeeklyWorkUsers[1].ScreenName
		workName3 = top3.WeeklyWorkUsers[2].ScreenName
		resultPoint1 = int(top3.WeeklyResultUsers[0].Point)
		resultPoint2 = int(top3.WeeklyResultUsers[1].Point)
		resultPoint3 = int(top3.WeeklyResultUsers[2].Point)
		resultName1 = top3.WeeklyResultUsers[0].ScreenName
		resultName2 = top3.WeeklyResultUsers[1].ScreenName
		resultName3 = top3.WeeklyResultUsers[2].ScreenName
	case "monthly":
		workPoint1 = top3.MonthlyWorkUsers[0].Point
		workPoint2 = top3.MonthlyWorkUsers[1].Point
		workPoint3 = top3.MonthlyWorkUsers[2].Point
		workName1 = top3.MonthlyWorkUsers[0].ScreenName
		workName2 = top3.MonthlyWorkUsers[1].ScreenName
		workName3 = top3.MonthlyWorkUsers[2].ScreenName
		resultPoint1 = int(top3.MonthlyResultUsers[0].Point)
		resultPoint2 = int(top3.MonthlyResultUsers[1].Point)
		resultPoint3 = int(top3.MonthlyResultUsers[2].Point)
		resultName1 = top3.MonthlyResultUsers[0].ScreenName
		resultName2 = top3.MonthlyResultUsers[1].ScreenName
		resultName3 = top3.MonthlyResultUsers[2].ScreenName
	}

	body := fmt.Sprintf(`
【%s速報 %v】

［作業ランキング］
🥇 %gpt @%s
🥈 %gpt @%s
🥉 %gpt @%s

［成果ランキング］
🥇 %dpt @%s
🥈 %dpt @%s
🥉 %dpt @%s

▼ 4位以下はコチラ ▼
https://influ-dojo.work
`,
		period[path], time.Now().Format("1/2 15:04"),
		workPoint1, workName1, workPoint2, workName2, workPoint3, workName3,
		resultPoint1, resultName1, resultPoint2, resultName2, resultPoint3, resultName3,
	)

	fmt.Printf("body: %+v", body)
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

			return xerrors.Errorf("failed to favorite: %w", err)
		}
	}

	return nil
}
