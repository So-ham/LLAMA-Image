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
