package image

import (
	"fmt"
	"os"
	"path/filepath"
)

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
