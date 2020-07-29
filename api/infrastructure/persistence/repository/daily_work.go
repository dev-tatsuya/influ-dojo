package repository

import (
	"influ-dojo/api/domain/apperr"
	domainModel "influ-dojo/api/domain/model"
	"influ-dojo/api/domain/repository"
	dataModel "influ-dojo/api/infrastructure/persistence/model"

	"github.com/jinzhu/gorm"
)

type dailyWork gormRepository

func NewDailyWork(db *gorm.DB) repository.DailyWork {
	return &dailyWork{db}
}

func (work *dailyWork) LoadTop3() ([]*domainModel.Work, error) {
	mdls := make([]*dataModel.DailyWork, 0)
	if err := work.DB.Order("point desc").Limit(3).Find(&mdls).Error; err != nil {
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

func (work *dailyWork) LoadByScreenName(screenName string) (*domainModel.Work, error) {
	mdl := new(dataModel.DailyWork)
	if err := work.DB.Where("screen_name = ?", screenName).First(mdl).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, apperr.ErrRecordNotFound
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

func (work *dailyWork) Save(entity *domainModel.Work) error {
	mdl := &dataModel.DailyWork{
		ScreenName:             entity.ScreenName,
		TweetsCount:            entity.TweetsCount,
		IncreaseTweetsCount:    &entity.IncreaseTweetsCount,
		FavoritesCount:         entity.FavoritesCount,
		IncreaseFavoritesCount: &entity.IncreaseFavoritesCount,
		Point:                  &entity.Point,
	}

	return work.DB.Where("screen_name = ?", mdl.ScreenName).Assign(*mdl).FirstOrCreate(mdl).Error
}
