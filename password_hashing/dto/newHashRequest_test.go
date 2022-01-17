package dto

//All tests should obey the Triple A format. Arrange, Act, Assert

import (
	"net/http"
	"testing"
)

func Test_should_return_error_when_password_is_less_than_8_characters_long(t *testing.T) {
	// Arrange
	request := NewHashRequest{
		PasswordString: "Ps2d",
	}

	// Act
	appError := request.Validate()

	// Assert
	if appError.Message != "Password must be at least 8 characters long" {
		t.Error("Invalid message while testing password shorter than 8 characters")
	}

	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing password length")
	}
}

func Test_should_return_error_when_password_is_more_than_50_characters_long(t *testing.T) {
	// Arrange
	request := NewHashRequest{
		PasswordString: "Password1234567891234567891234567891234566778912345678uih",
	}

	// Act
	appError := request.Validate()

	// Assert
	if appError.Message != "Password must be less than 50 characters long" {
		t.Error("Invalid message while password longer than 50 characters")
	}

	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing password length")
	}
}

func Test_should_return_error_when_password_does_not_have_a_capital_letter(t *testing.T) {
	// Arrange
	request := NewHashRequest{
		PasswordString: "password12",
	}

	// Act
	appError := request.Validate()

	// Assert
	if appError.Message != "Password must include a capital letter" {
		t.Error("Invalid message while testing password with capital letter")
	}

	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing password with no capital letter")
	}
}

func Test_should_return_error_when_error_is_not_nil_with_proper_password(t *testing.T) {
	// Arrange
	request := NewHashRequest{
		PasswordString: "Password12",
	}

	// Act
	appError := request.Validate()

	// Assert
	if appError != nil {
		t.Error("Invalid error after passing correctly formatted password")
	}
}

func Test_should_return_error_when_id_is_not_a_number(t *testing.T) {
	// Arrange
	request := NewHashRequest{
		PasswordString: "Password12",
	}

	// Act
	_, appError := request.ValidateId("asd")

	// Assert
	if appError.Message != "Please provide valid identifier. (Numbers only)" {
		t.Error("Invalid message while testing hash id with letters")
	}
}
