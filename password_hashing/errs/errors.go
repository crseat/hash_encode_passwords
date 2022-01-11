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
