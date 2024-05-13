package service

import (
	"context"

	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/dto"
)

type PhotoService interface {
	Upload(ctx context.Context, userId uint64, data *dto.PhotoUploadFormRequest) (*dto.PhotoUploadFormResponse, error)
}
