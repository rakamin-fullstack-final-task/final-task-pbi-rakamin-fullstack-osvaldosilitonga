package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/domain"
	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/dto"
	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/helpers"
	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/repository"
)

type userServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userServiceImpl{
		userRepo: ur,
	}
}

func (us *userServiceImpl) AddNewUser(ctx context.Context, data *dto.UserRegisterRequest) (domain.User, error) {
	// Prepare the data
	user := domain.User{
		Username: data.Username,
		Email:    data.Email,
	}

	// Encrypt Password
	hash, err := helpers.HashPassword(data.Password)
	if err != nil {
		return user, errors.New("500,error when hashing password")
	}

	user.Password = hash

	// Save to Repository
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	if err = us.userRepo.Save(c, &user); err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return user, errors.New("400,email already exists")
		}

		return user, errors.New("500,error when inserting to database")
	}

	return user, nil
}
