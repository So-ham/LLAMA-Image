package chat

import "chat2/internal/entities"

// Chat represents the use case for chat functionality

type chatInteractor struct{}

type Chat interface {
	LlamaChatRequest(requestData *entities.ChatReq) (*entities.LlamaResp, error)
}

// NewChat returns a new instance of Chat
func New() Chat {
	return &chatInteractor{}
}
