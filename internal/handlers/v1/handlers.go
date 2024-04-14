package v1

import (
	"chat2/internal/services"
	"net/http"
)

// handlerV1 represents the v1 handler
type handlerV1 struct {
	Service services.Service
}

// HandlerV1 has handlers list
type HandlerV1 interface {
	ChatHandler(w http.ResponseWriter, r *http.Request)
	UploadImageHandler(w http.ResponseWriter, r *http.Request)
	GetImageHandler(w http.ResponseWriter, r *http.Request)
}

// New creates an instance of handlers for V1
func New(s services.Service) HandlerV1 {
	return &handlerV1{Service: s}
}
