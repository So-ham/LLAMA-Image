package rest

import (
	"chat2/internal/handlers"

	"github.com/gorilla/mux"
)

// NewRouter returns a new router instance with configured routes
func NewRouter(h *handlers.Handler) *mux.Router {
	router := mux.NewRouter()

	// Chat endpoint
	router.HandleFunc("/chat", h.V1.ChatHandler).Methods("POST")

	// Upload image endpoint
	router.HandleFunc("/upload", h.V1.UploadImageHandler).Methods("POST")

	// Get image endpoint
	router.HandleFunc("/images/{userID}", h.V1.GetImageHandler).Methods("GET")

	return router
}
