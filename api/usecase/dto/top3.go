package dto

type Top3 struct {
	DailyWorkUsers     []*TopUser `json:"daily_work_users"`
	DailyResultUsers   []*TopUser `json:"daily_result_users"`
	WeeklyWorkUsers    []*TopUser `json:"weekly_work_users"`
	WeeklyResultUsers  []*TopUser `json:"weekly_result_users"`
	MonthlyWorkUsers   []*TopUser `json:"monthly_work_users"`
	MonthlyResultUsers []*TopUser `json:"monthly_result_users"`
}

type TopUser struct {
	ScreenName string  `json:"screen_name"`
	Point      float64 `json:"point"`
}
