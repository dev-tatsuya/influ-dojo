package repository

import (
	"influ-dojo/api/domain/apperr"
	domainModel "influ-dojo/api/domain/model"
	"influ-dojo/api/domain/repository"
	dataModel "influ-dojo/api/infrastructure/persistence/model"

	"github.com/jinzhu/gorm"
	"golang.org/x/xerrors"
)

type weeklyWork struct {
	GormRepository
}

func NewWeeklyWork(db *gorm.DB) repository.Work {
	return &weeklyWork{GormRepository{db}}
}

func (repo *weeklyWork) LoadOrderByRanking() ([]*domainModel.Work, error) {
	mdls := make([]*dataModel.WeeklyWork, 0)
	if err := repo.DB.Order("point desc").Find(&mdls).Error; err != nil {
		return nil, xerrors.Errorf("failed to load work ranking: %w", err)
	}

	entities := make([]*domainModel.Work, 0)
	for _, mdl := range mdls {
		entities = append(entities, mdl.MakeEntity())
	}

	return entities, nil
}

func (repo *weeklyWork) LoadTop3() ([]*domainModel.Work, error) {
	mdls := make([]*dataModel.WeeklyWork, 0)
	if err := repo.DB.Order("point desc").Limit(3).Find(&mdls).Error; err != nil {
		return nil, err
	}

	entities := make([]*domainModel.Work, 0)
	for _, mdl := range mdls {
		entities = append(entities, mdl.MakeEntity())
	}

	return entities, nil
}

func (repo *weeklyWork) LoadByScreenName(screenName string) (*domainModel.Work, error) {
	mdl := new(dataModel.WeeklyWork)
	if err := repo.DB.Where("screen_name = ?", screenName).First(mdl).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, apperr.ErrRecordNotFound
		}

		return nil, err
	}

	return mdl.MakeEntity(), nil
}

func (repo *weeklyWork) Save(entity *domainModel.Work) error {
	mdl := &dataModel.WeeklyWork{
		ScreenName:             entity.ScreenName,
		TweetsCount:            entity.TweetsCount,
		IncreaseTweetsCount:    &entity.IncreaseTweetsCount,
		FavoritesCount:         entity.FavoritesCount,
		IncreaseFavoritesCount: &entity.IncreaseFavoritesCount,
		Point:                  &entity.Point,
		Ranking:                entity.Ranking,
		LastRanking:            entity.LastRanking,
	}

	return repo.store(repo.DB, mdl)
}
