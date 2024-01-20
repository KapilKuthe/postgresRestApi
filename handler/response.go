// response.go
package handler

import (
	"encoding/json"
	"net/http"
)

// APIResponse represents the standardized JSON response format.
type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// RespondJSON sends a JSON response to the client.
func RespondJSON(w http.ResponseWriter, status int, data interface{}, message string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	var response APIResponse

	if err != nil {
		response = APIResponse{
			Status: status,
			Error:  err.Error(),
		}
	} else {
		response = APIResponse{
			Status:  status,
			Message: message,
			Data:    data,
		}
	}

	json.NewEncoder(w).Encode(response)
}
