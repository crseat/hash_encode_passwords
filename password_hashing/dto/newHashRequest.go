//Package dto implements the data transfer objects and validates the input
package dto

import (
	"password_hashing/errs"
	"unicode"
)

type NewHashRequest struct {
	PasswordString string
	Id             int64
}

//Validate the password that was passed in.
func (r NewHashRequest) Validate() *errs.AppError {
	if len(r.PasswordString) < 8 {
		return errs.NewValidationError("Password must be at least 8 characters long")
	}
	if len(r.PasswordString) > 50 {
		return errs.NewValidationError("Password must be less than 50 characters long")
	}
	if !UppercaseAndNumberPresent(r.PasswordString) {
		return errs.NewValidationError("Password must include a number and a capital letter")
	}
	return nil
}

//Validate the hash identifier that was passed in.
func (r NewHashRequest) ValidateId() *errs.AppError {
	return nil
}

//UppercaseAndNumberPresent takes in a password and iterates over it returning whether it contains an uppercase letter
//and a number.
func UppercaseAndNumberPresent(password string) bool {
	uppercase := false
	number := false
	for _, c := range password {
		if unicode.IsUpper(c) {
			uppercase = true
		}
		if unicode.IsNumber(c) {
			number = true
		}
		if uppercase && number {
			return true
		}
	}
	return false
}
