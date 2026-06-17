package response

// APIResponse
type APIResponse struct {
	Data    any    `json:"data,omitempty"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// Success
func Success(message string, data any) APIResponse {
	return APIResponse{
		Message: message,
		Data:    data,
	}
}

// Error
func Error(message string, err string) APIResponse {
	return APIResponse{
		Message: message,
		Error:   err,
	}
}