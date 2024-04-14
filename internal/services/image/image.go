package image

import (
	"chat2/internal/entities"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
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

func (ii *imageInteractor) UploadImage(userID int, file multipart.File, filename string) error {
	// Create a new file in the images directory
	f, err := os.Create(filepath.Join("images", filename))
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer f.Close()

	// Copy the file content to the new file
	_, err = io.Copy(f, file)
	if err != nil {
		return fmt.Errorf("failed to copy file content: %v", err)
	}

	// Store metadata in memory
	ii.images[userID] = entities.ImageMetadata{
		UserID:    userID,
		Filename:  filename,
		Timestamp: time.Now(),
	}

	return nil
}

func (ii *imageInteractor) GetImage(userID int) ([]byte, error) {
	// Retrieve filename from metadata
	imageMeta, ok := ii.images[userID]
	if !ok {
		return nil, fmt.Errorf("image not found for user ID: %d", userID)
	}
	filename := imageMeta.Filename

	// Read the image file
	filePath := filepath.Join("images", filename)
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open image file: %v", err)
	}

	return file, nil
}
