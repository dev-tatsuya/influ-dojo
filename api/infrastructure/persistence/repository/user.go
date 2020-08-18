package repository

import (
	"influ-dojo/api/domain/apperr"
	domainModel "influ-dojo/api/domain/model"
	"influ-dojo/api/domain/repository"
	dataModel "influ-dojo/api/infrastructure/persistence/model"

	"github.com/jinzhu/gorm"
	"golang.org/x/xerrors"
)

type user struct {
	GormRepository
}

func NewUser(db *gorm.DB) repository.User {
	return &user{GormRepository{db}}
}

func (repo *user) LoadByID(userID string) (*domainModel.User, error) {
	mdl := new(dataModel.User)
	if err := repo.DB.Where("user_id = ?", userID).First(mdl).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, apperr.ErrRecordNotFound
		}

		return nil, err
	}

	return &domainModel.User{
		UserID:       mdl.UserID,
		Name:         mdl.Name,
		ScreenName:   mdl.ScreenName,
		ProfileImage: mdl.ProfileImage,
	}, nil
}

func (repo *user) LoadByScreenName(screenName string) (*domainModel.User, error) {
	mdl := new(dataModel.User)
	if err := repo.DB.Where("screen_name = ?", screenName).First(mdl).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, apperr.ErrRecordNotFound
		}

		return nil, err
	}

	return &domainModel.User{
		UserID:       mdl.UserID,
		Name:         mdl.Name,
		ScreenName:   mdl.ScreenName,
		ProfileImage: mdl.ProfileImage,
	}, nil
}

func (repo *user) LoadIDs() ([]string, error) {
	mdls := make([]*dataModel.User, 0)
	if err := repo.DB.Find(&mdls).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return []string{}, nil
		}

		return nil, xerrors.Errorf("failed to load ids: %w", err)
	}

	IDs := make([]string, 0)
	for _, mdl := range mdls {
		IDs = append(IDs, mdl.UserID)
	}

	return IDs, nil
}

func (repo *user) Save(entity *domainModel.User) error {
	mdl := &dataModel.User{
		UserID:       entity.UserID,
		Name:         entity.Name,
		ScreenName:   entity.ScreenName,
		ProfileImage: entity.ProfileImage,
	}

	return repo.store(repo.DB, mdl)
}

func (repo *user) Delete(userID string) error {
	user, err := repo.LoadByID(userID)
	if err != nil {
		return err
	}

	return transaction(repo.DB, func(tx *gorm.DB) error {
		if err := tx.Delete(&dataModel.User{}, "user_id = ?", userID).Error; err != nil {
			return xerrors.Errorf("failed to delete user_id = %s: %w", userID, err)
		}

		if err := tx.Delete(&dataModel.DailyWork{}, "screen_name = ?", user.ScreenName).Error; err != nil {
			return xerrors.Errorf("failed to delete daily work for screen_name = %s: %w", user.ScreenName, err)
		}

		if err := tx.Delete(&dataModel.DailyResult{}, "screen_name = ?", user.ScreenName).Error; err != nil {
			return xerrors.Errorf("failed to delete daily result for screen_name = %s: %w", user.ScreenName, err)
		}

		if err := tx.Delete(&dataModel.WeeklyWork{}, "screen_name = ?", user.ScreenName).Error; err != nil {
			return xerrors.Errorf("failed to delete weekly work for screen_name = %s: %w", user.ScreenName, err)
		}

		if err := tx.Delete(&dataModel.WeeklyResult{}, "screen_name = ?", user.ScreenName).Error; err != nil {
			return xerrors.Errorf("failed to delete weekly result for screen_name = %s: %w", user.ScreenName, err)
		}

		if err := tx.Delete(&dataModel.MonthlyWork{}, "screen_name = ?", user.ScreenName).Error; err != nil {
			return xerrors.Errorf("failed to delete monthly work for screen_name = %s: %w", user.ScreenName, err)
		}

		if err := tx.Delete(&dataModel.MonthlyResult{}, "screen_name = ?", user.ScreenName).Error; err != nil {
			return xerrors.Errorf("failed to delete monthly result for screen_name = %s: %w", user.ScreenName, err)
		}

		if err := tx.Commit().Error; err != nil {
			return xerrors.Errorf("failed to commit delete user = %s: %w", user.ScreenName, err)
		}

		return nil
	})
}
