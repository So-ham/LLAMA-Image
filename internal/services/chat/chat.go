package chat

// Chat represents the use case for chat functionality

type chatInteractor struct{}

type Chat interface {
	LlamaChatRequest(requestData map[string]interface{}) (map[string]interface{}, error)
}

// NewChat returns a new instance of Chat
func New() Chat {
	return &chatInteractor{}
}
