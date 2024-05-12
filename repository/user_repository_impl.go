package repository

import (
	"context"

	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/domain"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		DB: db,
	}
}

func (ur *userRepositoryImpl) Save(ctx context.Context, user *domain.User) error {

	if err := ur.DB.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}

	return nil
}
