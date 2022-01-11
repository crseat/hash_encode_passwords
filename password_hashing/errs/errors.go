//Package errs defines and organizes error types and functions
package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}
