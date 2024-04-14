package entities

import "time"

// Image represents an uploaded image entity
type Image struct {
	UserID   int
	Filename string
}

// ImageMetadata stores information about uploaded images
type ImageMetadata struct {
	UserID    int
	Filename  string
	Timestamp time.Time
}
