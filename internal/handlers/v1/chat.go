package v1

import (
	"chat2/internal/entities"
	"encoding/json"
	"net/http"
)

func (h *handlerV1) ChatHandler(w http.ResponseWriter, r *http.Request) {

	// Read user messages from request body
	var requestData *entities.ChatReq
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Handle chat request
	responseData, err := h.Service.Chat.LlamaChatRequest(requestData)
	if err != nil {
		http.Error(w, "Failed to handle chat request", http.StatusInternalServerError)
		return
	}

	// Forward response to client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}
