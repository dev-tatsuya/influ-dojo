package dto

type RankingAll struct {
	DailyWorkRanking     *Ranking `json:"daily_work_ranking"`
	DailyResultRanking   *Ranking `json:"daily_result_ranking"`
	WeeklyWorkRanking    *Ranking `json:"weekly_work_ranking"`
	WeeklyResultRanking  *Ranking `json:"weekly_result_ranking"`
	MonthlyWorkRanking   *Ranking `json:"monthly_work_ranking"`
	MonthlyResultRanking *Ranking `json:"monthly_result_ranking"`
}

type Ranking struct {
	RankUsers []*RankUser `json:"rank_users"`
}

type RankUser struct {
	Name         string  `json:"name"`
	ScreenName   string  `json:"screen_name"`
	ProfileImage string  `json:"profile_image"`
	Point        float64 `json:"point"`
	Ranking      int     `json:"ranking"`
	LastRanking  int     `json:"last_ranking"`
}
