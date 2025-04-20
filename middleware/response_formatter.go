package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSONResponse(w http.ResponseWriter, statusCode int, success bool, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := Response{
		Success: success,
		Message: message,
		Data:    data,
	}
	log.Printf("[%s] Response: %d %s", time.Now().Format(time.RFC3339), statusCode, message)
	json.NewEncoder(w).Encode(response)
}

func JSONErrorResponse(w http.ResponseWriter, statusCode int, message string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := Response{
		Success: false,
		Message: message,
		Data:    map[string]string{"error": err.Error()},
	}
	log.Printf("[%s] Error Response: %d %s - %v", time.Now().Format(time.RFC3339), statusCode, message, err)
	json.NewEncoder(w).Encode(response)
}