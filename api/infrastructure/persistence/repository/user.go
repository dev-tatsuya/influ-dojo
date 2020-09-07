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
	return transaction(repo.DB, func(tx *gorm.DB) error {
		if err := tx.Unscoped().Delete(&dataModel.User{}, "user_id = ?", userID).Error; err != nil {
			return xerrors.Errorf("failed to delete user_id = %s: %w", userID, err)
		}

		if err := tx.Unscoped().Delete(&dataModel.DailyWork{}, "user_id = ?", userID).Error; err != nil {
			return xerrors.Errorf("failed to delete daily work for user_id = %s: %w", userID, err)
		}

		if err := tx.Unscoped().Delete(&dataModel.DailyResult{}, "user_id = ?", userID).Error; err != nil {
			return xerrors.Errorf("failed to delete daily result for user_id = %s: %w", userID, err)
		}

		if err := tx.Unscoped().Delete(&dataModel.WeeklyWork{}, "user_id = ?", userID).Error; err != nil {
			return xerrors.Errorf("failed to delete weekly work for user_id = %s: %w", userID, err)
		}

		if err := tx.Unscoped().Delete(&dataModel.WeeklyResult{}, "user_id = ?", userID).Error; err != nil {
			return xerrors.Errorf("failed to delete weekly result for user_id = %s: %w", userID, err)
		}

		if err := tx.Unscoped().Delete(&dataModel.MonthlyWork{}, "user_id = ?", userID).Error; err != nil {
			return xerrors.Errorf("failed to delete monthly work for user_id = %s: %w", userID, err)
		}

		if err := tx.Unscoped().Delete(&dataModel.MonthlyResult{}, "user_id = ?", userID).Error; err != nil {
			return xerrors.Errorf("failed to delete monthly result for user_id = %s: %w", userID, err)
		}

		if err := tx.Commit().Error; err != nil {
			return xerrors.Errorf("failed to commit delete user = %s: %w", userID, err)
		}

		return nil
	})
}
