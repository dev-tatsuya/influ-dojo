package output

type DailyRank struct {
	WorkRank   *WorkRank   `json:"work_rank"`
	ResultRank *ResultRank `json:"result_rank"`
}

type WorkRank struct {
	WorkUsers []*WorkUser `json:"work_users"`
}

type ResultRank struct {
	ResultUsers []*ResultUser `json:"result_users"`
}

type WorkUser struct {
	Name                   string `json:"name"`
	ScreenName             string `json:"screen_name"`
	ProfileImage           string `json:"profile_image"`
	IncreaseTweetsCount    int    `json:"increase_tweets_count"`
	IncreaseFavoritesCount int    `json:"increase_favorites_count"`
}

type ResultUser struct {
	Name                   string `json:"name"`
	ScreenName             string `json:"screen_name"`
	ProfileImage           string `json:"profile_image"`
	IncreaseFollowersCount int    `json:"increase_followers_count"`
}
