package services

import (
	"chat2/internal/models"
	"chat2/internal/services/chat"
	"chat2/internal/services/image"
)

// Service represents the service layer having
// all the services from all service packages
type Service struct {
	Chat  chat.Chat
	Image image.Image
}

// New creates a new instance of Service
func New(model *models.Model) Service {
	s := Service{
		Chat:  chat.New(),
		Image: image.New(),
	}
	return s
}
