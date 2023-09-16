package dto

type ApiResponse[T any] struct {
	StatusCode int    `json:"status_code,omitempty"`
	Status     string `json:"status,omitempty"`
	Message    string `json:"message,omitempty"`
	Data       T      `json:"data,omitempty"`
}
