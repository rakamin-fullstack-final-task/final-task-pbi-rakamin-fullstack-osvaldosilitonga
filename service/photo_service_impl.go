package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/domain"
	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/dto"
	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/repository"
)

type photoServiceImpl struct {
	PhotoRepo repository.PhotoRepository
}

func NewPhotoService(pr repository.PhotoRepository) PhotoService {
	return &photoServiceImpl{
		PhotoRepo: pr,
	}
}

func (ps *photoServiceImpl) Upload(ctx context.Context, userID uint64, data *dto.PhotoUploadFormRequest) (*dto.PhotoUploadFormResponse, error) {
	var res dto.PhotoUploadFormResponse

	photo := domain.Photo{
		UserID:   userID,
		Title:    data.Title,
		Caption:  data.Caption,
		PhotoUrl: data.Path,
	}

	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	ph, err := ps.PhotoRepo.Save(c, &photo)
	if err != nil {
		if strings.Contains(err.Error(), "foreign key constraint") {
			return &res, errors.New("404,user id not exist")
		}

		return &res, errors.New(fmt.Sprintf("500,%v", err.Error()))
	}

	res.ID = ph.ID
	res.Title = ph.Title
	res.Caption = ph.Caption
	res.CreatedAt = ph.CreatedAt.Local().String()
	res.UpdatedAt = ph.UpdatedAt.Local().String()

	return &res, nil
}
