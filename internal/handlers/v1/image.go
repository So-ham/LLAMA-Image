package v1

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (h *handlerV1) UploadImageHandler(w http.ResponseWriter, r *http.Request) {

	// Parse form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get user ID from form
	userID, err := strconv.Atoi(r.FormValue("userID"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Get the file from the request
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Generate a unique filename
	filename := fmt.Sprintf("%d_%d.jpg", userID, time.Now().Unix())

	// Handle image upload
	err = h.Service.Image.UploadImage(userID, file, filename)
	if err != nil {
		http.Error(w, "Failed to handle image upload", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Image uploaded successfully"))
}

func (h *handlerV1) GetImageHandler(w http.ResponseWriter, r *http.Request) {

	// Get user ID from request URL
	userID, err := strconv.Atoi(mux.Vars(r)["userID"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Handle image retrieval
	file, err := h.Service.Image.GetImage(userID)
	if err != nil {
		http.Error(w, "Failed to retrieve image", http.StatusInternalServerError)
		return
	}

	// Set Content-Type header
	w.Header().Set("Content-Type", "image/jpeg")

	// Write image file to response body
	w.Write(file)
	fmt.Println("Image retrieved successfully")
}
