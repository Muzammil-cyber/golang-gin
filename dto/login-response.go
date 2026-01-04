package dto

import "github.com/muzammil-cyber/golang-gin/utils"

// LoginResponse represents the login response with JWT token
type LoginResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."` // JWT token
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error" example:"Invalid credentials"` // Error message
}

// ValidationErrorResponse represents validation error response
type ValidationErrorResponse struct {
	Errors []utils.ValidationError `json:"errors"` // Validation errors as string
}
