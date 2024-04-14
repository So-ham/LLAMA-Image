package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func (ci *chatInteractor) LlamaChatRequest(requestData map[string]interface{}) (map[string]interface{}, error) {
	// Prepare data to send to Python server
	requestDataBytes, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request data: %v", err)
	}

	// Send data to LLAMA server
	LlamaServerURL := os.Getenv("LLAMA_SERVER_URL")
	fmt.Println(LlamaServerURL)
	llamaResp, err := http.Post(LlamaServerURL, "application/json", bytes.NewBuffer(requestDataBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to send request to Python server: %v", err)
	}
	defer llamaResp.Body.Close()

	// Read Python server response
	var responseData map[string]interface{}
	err = json.NewDecoder(llamaResp.Body).Decode(&responseData)
	if err != nil {
		return nil, fmt.Errorf("failed to read response from Python server: %v", err)
	}

	return responseData, nil
}
