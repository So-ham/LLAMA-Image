package image

import (
	"chat2/internal/entities"
	"mime/multipart"
)

type Image interface {
	UploadImage(userID int, file multipart.File, filename string) error
	GetImage(userID int) ([]byte, error)
}

type imageInteractor struct {
	images map[int]entities.ImageMetadata
}

// NewImage returns a new instance of Image
func New() Image {
	return &imageInteractor{
		images: make(map[int]entities.ImageMetadata),
	}
}
