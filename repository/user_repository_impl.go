package repository

import (
	"context"

	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (ur *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	user := domain.User{}

	err := ur.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error

	return user, err
}

func (ur *userRepositoryImpl) FindByID(ctx context.Context, id uint64) (*domain.User, error) {
	user := domain.User{ID: id}

	err := ur.DB.First(&user).Error

	return &user, err
}

func (ur *userRepositoryImpl) Update(ctx context.Context, id uint64, username, password string) (*domain.User, error) {
	var user domain.User

	query := map[string]any{}

	if username != "" {
		query["username"] = username
	}
	if password != "" {
		query["password"] = password
	}

	err := ur.DB.Model(&user).Clauses(clause.Returning{}).Limit(1).Where("id = ?", id).Updates(query).Error

	return &user, err
}

func (ur *userRepositoryImpl) DeleteByID(ctx context.Context, id uint64) error {
	if _, err := ur.FindByID(ctx, id); err != nil {
		return err
	}

	err := ur.DB.Delete(domain.User{}, id).Error

	return err
}
