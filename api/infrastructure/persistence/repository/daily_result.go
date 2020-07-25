package repository

import (
	"errors"
	domainModel "influ-dojo/api/domain/model"
	"influ-dojo/api/domain/repository"
	dataModel "influ-dojo/api/infrastructure/persistence/model"

	"github.com/jinzhu/gorm"
)

type dailyResult gormRepository

func NewDailyResult(db *gorm.DB) repository.DailyResult {
	return &dailyResult{db}
}

func (dw *dailyResult) Load() ([]*domainModel.Result, error) {
	mdls := make([]*dataModel.DailyResult, 0)
	if err := dw.DB.Find(&mdls).Error; err != nil {
		return nil, err
	}

	entities := make([]*domainModel.Result, len(mdls))
	for _, mdl := range mdls {
		entities = append(entities, &domainModel.Result{
			UserID:         mdl.UserID,
			FollowersCount: mdl.FollowersCount,
		})
	}

	return entities, nil
}

func (dw *dailyResult) LoadByID(userID string) (*domainModel.Result, error) {
	mdl := new(dataModel.DailyResult)
	if err := dw.DB.Where("user_id = ?", userID).First(mdl).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("not found")
		}

		return nil, err
	}

	return &domainModel.Result{
		UserID:         mdl.UserID,
		FollowersCount: mdl.FollowersCount,
	}, nil
}

func (dw *dailyResult) Save(entity *domainModel.Result) error {
	mdl := &dataModel.DailyResult{
		UserID:         entity.UserID,
		FollowersCount: entity.FollowersCount,
	}

	return dw.DB.Where("user_id = ?", mdl.UserID).Assign(mdl).FirstOrCreate(mdl).Error
}
