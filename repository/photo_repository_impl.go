package repository

import (
	"context"

	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/domain"
	"gorm.io/gorm"
)

type photoRepositoryImpl struct {
	DB *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoRepositoryImpl{
		DB: db,
	}
}

func (pr *photoRepositoryImpl) Save(ctx context.Context, photo *domain.Photo) (*domain.Photo, error) {
	err := pr.DB.Save(&photo).Error

	return photo, err
}
