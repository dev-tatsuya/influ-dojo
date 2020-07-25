package repository

import "github.com/jinzhu/gorm"

type gormRepository struct {
	DB *gorm.DB
}
