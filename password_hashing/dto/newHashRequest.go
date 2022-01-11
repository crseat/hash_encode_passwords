package dto

import (
	"password_hashing/errs"
	"unicode"
)

type NewHashRequest struct {
	password string
}

//Validate input from user
func (r NewHashRequest) Validate() *errs.AppError {
	if len(r.password) < 8 {
		return errs.NewValidationError("Password must be at least 8 characters long")
	}
	if len(r.password) > 50 {
		return errs.NewValidationError("Password must be less than 50 characters long")
	}
	if !UppercaseAndNumberPresent(r.password) {
		return errs.NewValidationError("Password must include a number and a capital letter")
	}
	return nil
}

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
