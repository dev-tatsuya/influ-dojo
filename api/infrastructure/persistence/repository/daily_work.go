package repository

import (
	"errors"
	domainModel "influ-dojo/api/domain/model"
	"influ-dojo/api/domain/repository"
	dataModel "influ-dojo/api/infrastructure/persistence/model"

	"github.com/jinzhu/gorm"
)

type dailyWork gormRepository

func NewDailyWork(db *gorm.DB) repository.DailyWork {
	return &dailyWork{db}
}

func (dw *dailyWork) Load() ([]*domainModel.Work, error) {
	mdls := make([]*dataModel.DailyWork, 0)
	if err := dw.DB.Find(&mdls).Error; err != nil {
		return nil, err
	}

	entities := make([]*domainModel.Work, len(mdls))
	for _, mdl := range mdls {
		entities = append(entities, &domainModel.Work{
			UserID:         mdl.UserID,
			TweetsCount:    mdl.TweetsCount,
			FavoritesCount: mdl.FavoritesCount,
		})
	}

	return entities, nil
}

func (dw *dailyWork) LoadByID(userID string) (*domainModel.Work, error) {
	mdl := new(dataModel.DailyWork)
	if err := dw.DB.Where("user_id = ?", userID).First(mdl).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("not found")
		}

		return nil, err
	}

	return &domainModel.Work{
		UserID:         mdl.UserID,
		TweetsCount:    mdl.TweetsCount,
		FavoritesCount: mdl.FavoritesCount,
	}, nil
}

func (dw *dailyWork) Save(entity *domainModel.Work) error {
	mdl := &dataModel.DailyWork{
		UserID:         entity.UserID,
		TweetsCount:    entity.TweetsCount,
		FavoritesCount: entity.FavoritesCount,
	}

	return dw.DB.Where("user_id = ?", mdl.UserID).Assign(mdl).FirstOrCreate(mdl).Error
}
