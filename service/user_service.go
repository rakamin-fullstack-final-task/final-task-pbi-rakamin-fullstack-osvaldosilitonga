package service

import (
	"context"

	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/domain"
	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/dto"
)

type UserService interface {
	AddNewUser(ctx context.Context, user *dto.UserRegisterRequest) (domain.User, error)
	Login(ctx context.Context, data *dto.UserLoginRequest) (dto.UserLoginResponse, error)
	Update(ctx context.Context, data *dto.UserUpdateRequest) (*dto.UserUpdateResponse, error)
	DeleteByID(ctx context.Context, id uint64) error
}
