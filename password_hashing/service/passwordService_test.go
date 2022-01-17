package service

import (
	"password_hashing/dto"
	"testing"
	"time"
)

func Test_should_return_a_validation_error_response_when_the_request_is_not_validated(t *testing.T) {
	// Arrange
	request := dto.NewHashRequest{
		PasswordString: "test",
	}
	service := NewPasswordService(nil)
	// Act
	var test = time.Time{}
	_, appError := service.NewHash(request, test, nil)
	// Assert
	if appError == nil {
		t.Error("failed while testing the new hash validation")
	}
}
