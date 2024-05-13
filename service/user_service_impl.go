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
	"gorm.io/gorm"
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

func (us *userServiceImpl) Login(ctx context.Context, data *dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	res := dto.UserLoginResponse{}

	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	user, err := us.userRepo.FindByEmail(c, data.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.New("404,email not exist")
		}
		return res, errors.New("500,something went wrong")
	}

	// Compare password
	if !helpers.ComparePassword(user.Password, data.Password) {
		return res, errors.New("400,wrong password")
	}

	// Generate token
	token, err := helpers.GenerateToken(user.ID, user.Email)
	if err != nil {
		return res, errors.New("500,something went wrong")
	}

	res.AccessToken = token
	return res, nil
}

func (us *userServiceImpl) Update(ctx context.Context, data *dto.UserUpdateRequest) (*dto.UserUpdateResponse, error) {
	var res dto.UserUpdateResponse

	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	if data.Password != "" {
		hash, err := helpers.HashPassword(data.Password)
		if err != nil {
			return &res, errors.New("500,something went wrong")
		}

		data.Password = hash
	}

	id := ctx.Value("id").(uint64)

	user, err := us.userRepo.Update(c, id, data.Username, data.Password)
	if err != nil {
		return &res, errors.New("500,something went wrong")
	}

	res.ID = user.ID
	res.Username = user.Username
	res.Email = user.Email
	res.CreatedAt = user.CreatedAt.String()
	res.UpdatedAt = user.UpdatedAt.String()

	return &res, nil
}
