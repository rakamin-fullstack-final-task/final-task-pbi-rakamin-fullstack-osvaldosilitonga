package repository

import (
	"context"

	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	FindByID(ctx context.Context, id uint64) (*domain.User, error)
	Update(ctx context.Context, id uint64, username, password string) (*domain.User, error)
}
