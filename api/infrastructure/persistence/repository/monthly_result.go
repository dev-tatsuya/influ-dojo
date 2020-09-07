package repository

import (
	"errors"
	domainModel "influ-dojo/api/domain/model"
	"influ-dojo/api/domain/repository"
	dataModel "influ-dojo/api/infrastructure/persistence/model"

	"github.com/jinzhu/gorm"
	"golang.org/x/xerrors"
)

type monthlyResult struct {
	GormRepository
}

func NewMonthlyResult(db *gorm.DB) repository.Result {
	return &monthlyResult{GormRepository{db}}
}

func (repo *monthlyResult) LoadOrderByRanking() ([]*domainModel.Result, error) {
	mdls := make([]*dataModel.MonthlyResult, 0)
	if err := repo.DB.Order("point desc").Find(&mdls).Error; err != nil {
		return nil, xerrors.Errorf("failed to load result ranking: %w", err)
	}

	entities := make([]*domainModel.Result, 0)
	for _, mdl := range mdls {
		entities = append(entities, mdl.MakeEntity())
	}

	return entities, nil
}

func (repo *monthlyResult) LoadByID(id string) (*domainModel.Result, error) {
	mdl := new(dataModel.MonthlyResult)
	if err := repo.DB.Where("user_id = ?", id).First(mdl).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("not found")
		}

		return nil, err
	}

	return mdl.MakeEntity(), nil
}

func (repo *monthlyResult) Save(entity *domainModel.Result) error {
	mdl := &dataModel.MonthlyResult{
		UserID:                 entity.UserID,
		FollowersCount:         entity.FollowersCount,
		IncreaseFollowersCount: &entity.IncreaseFollowersCount,
		Point:                  &entity.Point,
		Ranking:                entity.Ranking,
		LastRanking:            entity.LastRanking,
	}

	return repo.store(repo.DB, mdl)
}
