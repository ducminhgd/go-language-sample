package response

import (
	"encoding/json"
	"net/http"
)

// Response is the standard API response structure
type Response struct {
	Error int         `json:"error"`
	Data  interface{} `json:"data"`
}

// TranslationResponse is the response structure for translation requests
type TranslationResponse struct {
	Message string `json:"message"`
}

// NewSuccessResponse creates a new success response
func NewSuccessResponse(data interface{}) Response {
	return Response{
		Error: 0,
		Data:  data,
	}
}

// NewErrorResponse creates a new error response
func NewErrorResponse(data interface{}) Response {
	return Response{
		Error: 1,
		Data:  data,
	}
}

// JSON writes the response as JSON to the http.ResponseWriter
func JSON(w http.ResponseWriter, statusCode int, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
