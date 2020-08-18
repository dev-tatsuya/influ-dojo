package repository

import (
	"influ-dojo/api/infrastructure/persistence/model"
	"log"

	"github.com/jinzhu/gorm"
	"golang.org/x/xerrors"
)

type GormRepository struct {
	DB *gorm.DB
}

func (repo *GormRepository) store(db *gorm.DB, mdl model.GormModel) error {
	if mdl.IsNew() {
		if err := mdl.AttachID(); err != nil {
			return xerrors.Errorf("failed to attach ID: %w", err)
		}

		if err := db.Create(mdl).Error; err != nil {
			return xerrors.Errorf("failed to create new record: %w", err)
		}
	} else {
		if err := db.Save(mdl).Error; err != nil {
			return xerrors.Errorf("failed to save record: %w", err)
		}
	}

	return nil
}

func transaction(db *gorm.DB, f func(tx *gorm.DB) error) error {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return xerrors.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		rollback(tx)

		if r := recover(); r != nil {
			panic(r)
		}
	}()

	return f(tx)
}

func rollback(tx *gorm.DB) {
	if err := tx.Rollback().Error; err != nil {
		log.Printf("failed to rollback transaction: %v", err)
		return
	}
}
