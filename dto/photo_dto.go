package dto

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type PhotoUploadFormRequest struct {
	UUID    uuid.UUID
	Title   string                `form:"title" binding:"required,min=3"`
	Caption string                `form:"caption" binding:"required,min=3"`
	Image   *multipart.FileHeader `form:"image" binding:"required"`
	Path    string
}

type PhotoUploadFormResponse struct {
	ID        uint64 `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
