package repository

import (
	"errors"
	domainModel "influ-dojo/api/domain/model"
	"influ-dojo/api/domain/repository"
	dataModel "influ-dojo/api/infrastructure/persistence/model"

	"github.com/jinzhu/gorm"
	"golang.org/x/xerrors"
)

type weeklyResult struct {
	gormRepository
}

func NewWeeklyResult(db *gorm.DB) repository.Result {
	return &weeklyResult{gormRepository{db}}
}

func (repo *weeklyResult) LoadOrderByRanking() ([]*domainModel.Result, error) {
	mdls := make([]*dataModel.WeeklyResult, 0)
	if err := repo.DB.Order("point desc").Find(&mdls).Error; err != nil {
		return nil, xerrors.Errorf("failed to load result ranking: %w", err)
	}

	entities := make([]*domainModel.Result, 0)
	for _, mdl := range mdls {
		entities = append(entities, mdl.MakeEntity())
	}

	return entities, nil
}

func (repo *weeklyResult) LoadTop3() ([]*domainModel.Result, error) {
	mdls := make([]*dataModel.WeeklyResult, 0)
	if err := repo.DB.Order("point desc").Limit(3).Find(&mdls).Error; err != nil {
		return nil, err
	}

	entities := make([]*domainModel.Result, 0)
	for _, mdl := range mdls {
		entities = append(entities, mdl.MakeEntity())
	}

	return entities, nil
}

func (repo *weeklyResult) LoadByScreenName(screenName string) (*domainModel.Result, error) {
	mdl := new(dataModel.WeeklyResult)
	if err := repo.DB.Where("screen_name = ?", screenName).First(mdl).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("not found")
		}

		return nil, err
	}

	return mdl.MakeEntity(), nil
}

func (repo *weeklyResult) Save(entity *domainModel.Result) error {
	mdl := &dataModel.WeeklyResult{
		ScreenName:             entity.ScreenName,
		FollowersCount:         entity.FollowersCount,
		IncreaseFollowersCount: &entity.IncreaseFollowersCount,
		Point:                  &entity.Point,
		Ranking:                entity.Ranking,
		LastRanking:            entity.LastRanking,
	}

	return repo.store(repo.DB, mdl)
}
