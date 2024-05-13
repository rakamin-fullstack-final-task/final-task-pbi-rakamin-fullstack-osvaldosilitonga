package repository

import (
	"context"

	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/domain"
)

type PhotoRepository interface {
	Save(ctx context.Context, photo *domain.Photo) (*domain.Photo, error)
}
