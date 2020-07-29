package repository

import (
	"errors"
	domainModel "influ-dojo/api/domain/model"
	"influ-dojo/api/domain/repository"
	dataModel "influ-dojo/api/infrastructure/persistence/model"

	"github.com/jinzhu/gorm"
)

type monthlyWork gormRepository

func NewMonthlyWork(db *gorm.DB) repository.MonthlyWork {
	return &monthlyWork{db}
}

func (dw *monthlyWork) LoadTop3() ([]*domainModel.Work, error) {
	mdls := make([]*dataModel.MonthlyWork, 0)
	if err := dw.DB.Order("point desc").Limit(3).Find(&mdls).Error; err != nil {
		return nil, err
	}

	count := 0
	entities := make([]*domainModel.Work, len(mdls))
	for _, mdl := range mdls {
		if mdl.IncreaseTweetsCount == nil {
			mdl.IncreaseTweetsCount = &count
		}
		if mdl.IncreaseFavoritesCount == nil {
			mdl.IncreaseFavoritesCount = &count
		}
		if mdl.Point == nil {
			mdl.Point = &count
		}

		entities = append(entities, &domainModel.Work{
			ScreenName:             mdl.ScreenName,
			TweetsCount:            mdl.TweetsCount,
			IncreaseTweetsCount:    *mdl.IncreaseTweetsCount,
			FavoritesCount:         mdl.FavoritesCount,
			IncreaseFavoritesCount: *mdl.IncreaseFavoritesCount,
			Point:                  *mdl.Point,
		})
	}

	return entities, nil
}

func (dw *monthlyWork) LoadByScreenName(screenName string) (*domainModel.Work, error) {
	mdl := new(dataModel.MonthlyWork)
	if err := dw.DB.Where("screen_name = ?", screenName).First(mdl).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("not found")
		}

		return nil, err
	}

	count := 0
	if mdl.IncreaseTweetsCount == nil {
		mdl.IncreaseTweetsCount = &count
	}
	if mdl.IncreaseFavoritesCount == nil {
		mdl.IncreaseFavoritesCount = &count
	}
	if mdl.Point == nil {
		mdl.Point = &count
	}

	return &domainModel.Work{
		ScreenName:             mdl.ScreenName,
		TweetsCount:            mdl.TweetsCount,
		IncreaseTweetsCount:    *mdl.IncreaseTweetsCount,
		FavoritesCount:         mdl.FavoritesCount,
		IncreaseFavoritesCount: *mdl.IncreaseFavoritesCount,
		Point:                  *mdl.Point,
	}, nil
}

func (dw *monthlyWork) Save(entity *domainModel.Work) error {
	mdl := &dataModel.MonthlyWork{
		ScreenName:             entity.ScreenName,
		TweetsCount:            entity.TweetsCount,
		IncreaseTweetsCount:    &entity.IncreaseTweetsCount,
		FavoritesCount:         entity.FavoritesCount,
		IncreaseFavoritesCount: &entity.IncreaseFavoritesCount,
		Point:                  &entity.Point,
	}

	return dw.DB.Where("screen_name = ?", mdl.ScreenName).Assign(*mdl).FirstOrCreate(mdl).Error
}
