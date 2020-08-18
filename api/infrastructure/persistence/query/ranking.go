package queryService

import (
	"influ-dojo/api/infrastructure/persistence/repository"
	"influ-dojo/api/usecase/dto"

	"github.com/jinzhu/gorm"
)

type ranking struct {
	repository.GormRepository
}

func NewRanking(db *gorm.DB) *ranking {
	return &ranking{repository.GormRepository{DB: db}}
}

func (query *ranking) LoadRankingAll() (*dto.RankingAll, error) {
	dailyWorkRankUsers := make([]*dto.RankUser, 0)
	if err := query.DB.Table("daily_works").Order("point desc").
		Select("users.name, users.screen_name, users.profile_image, users.deleted_at, daily_works.point, daily_works.ranking, daily_works.last_ranking").
		Joins("INNER JOIN users ON daily_works.screen_name = users.screen_name AND users.deleted_at IS NULL").
		Scan(&dailyWorkRankUsers).Error; err != nil {
		return nil, err
	}

	dailyResultRankUsers := make([]*dto.RankUser, 0)
	if err := query.DB.Table("daily_results").Order("point desc").
		Select("users.name, users.screen_name, users.profile_image, users.deleted_at, daily_results.point, daily_results.ranking, daily_results.last_ranking").
		Joins("INNER JOIN users ON daily_results.screen_name = users.screen_name AND users.deleted_at IS NULL").
		Scan(&dailyResultRankUsers).Error; err != nil {
		return nil, err
	}

	weeklyWorkRankUsers := make([]*dto.RankUser, 0)
	if err := query.DB.Table("weekly_works").Order("point desc").
		Select("users.name, users.screen_name, users.profile_image, users.deleted_at, weekly_works.point, weekly_works.ranking, weekly_works.last_ranking").
		Joins("INNER JOIN users ON weekly_works.screen_name = users.screen_name AND users.deleted_at IS NULL").
		Scan(&weeklyWorkRankUsers).Error; err != nil {
		return nil, err
	}

	weeklyResultRankUsers := make([]*dto.RankUser, 0)
	if err := query.DB.Table("weekly_results").Order("point desc").
		Select("users.name, users.screen_name, users.profile_image, users.deleted_at, weekly_results.point, weekly_results.ranking, weekly_results.last_ranking").
		Joins("INNER JOIN users ON weekly_results.screen_name = users.screen_name AND users.deleted_at IS NULL").
		Scan(&weeklyResultRankUsers).Error; err != nil {
		return nil, err
	}

	monthlyWorkRankUsers := make([]*dto.RankUser, 0)
	if err := query.DB.Table("monthly_works").Order("point desc").
		Select("users.name, users.screen_name, users.profile_image, users.deleted_at, monthly_works.point, monthly_works.ranking, monthly_works.last_ranking").
		Joins("INNER JOIN users ON monthly_works.screen_name = users.screen_name AND users.deleted_at IS NULL").
		Scan(&monthlyWorkRankUsers).Error; err != nil {
		return nil, err
	}

	monthlyResultRankUsers := make([]*dto.RankUser, 0)
	if err := query.DB.Table("monthly_results").Order("point desc").
		Select("users.name, users.screen_name, users.profile_image, users.deleted_at, monthly_results.point, monthly_results.ranking, monthly_results.last_ranking").
		Joins("INNER JOIN users ON monthly_results.screen_name = users.screen_name AND users.deleted_at IS NULL").
		Scan(&monthlyResultRankUsers).Error; err != nil {
		return nil, err
	}

	return &dto.RankingAll{
		DailyWorkRanking:     &dto.Ranking{RankUsers: dailyWorkRankUsers},
		DailyResultRanking:   &dto.Ranking{RankUsers: dailyResultRankUsers},
		WeeklyWorkRanking:    &dto.Ranking{RankUsers: weeklyWorkRankUsers},
		WeeklyResultRanking:  &dto.Ranking{RankUsers: weeklyResultRankUsers},
		MonthlyWorkRanking:   &dto.Ranking{RankUsers: monthlyWorkRankUsers},
		MonthlyResultRanking: &dto.Ranking{RankUsers: monthlyResultRankUsers},
	}, nil
}
