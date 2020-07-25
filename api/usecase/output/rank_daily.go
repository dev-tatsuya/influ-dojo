package output

type DailyRank struct {
	WorkRank   *WorkRank   `json:"work_rank"`
	ResultRank *resultRank `json:"result_rank"`
}

type WorkRank struct {
	WorkUsers []*WorkUser `json:"work_users"`
}

type resultRank struct {
	ResultUsers []*resultUser `json:"result_users"`
}

type WorkUser struct {
	Name                   string `json:"name"`
	ScreenName             string `json:"screen_name"`
	ProfileImage           string `json:"profile_image"`
	IncreaseTweetsCount    int    `json:"increase_tweets_count"`
	IncreaseFavoritesCount int    `json:"increase_favorites_count"`
}

type resultUser struct {
	Name                   string `json:"name"`
	ScreenName             string `json:"screen_name"`
	ProfileImage           string `json:"profile_image"`
	IncreaseFollowersCount int    `json:"increase_followers_count"`
}
