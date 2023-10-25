package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

type APIResponse struct {
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
	Status    int         `json:"status"`
	Timestamp string      `json:"timestamp"`
}

func JsonResponse(w http.ResponseWriter, data interface{}, message string, status int) {
	response := APIResponse{
		Data:      data,
		Message:   message,
		Status:    status,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Write the JSON data to the response
	w.WriteHeader(status)
	w.Write(jsonResponse)
}
