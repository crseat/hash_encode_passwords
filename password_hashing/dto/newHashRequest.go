//Package dto implements the data transfer objects and validates the input
package dto

import (
	"password_hashing/errs"
	"strconv"
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
	if !UppercasePresent(r.PasswordString) {
		return errs.NewValidationError("Password must include a capital letter")
	}
	return nil
}

//Validate the hash identifier that was passed in.
func (r NewHashRequest) ValidateId(id string) (int, *errs.AppError) {
	hashId, err := strconv.Atoi(id)
	if err != nil {
		appError := errs.NewValidationError("Please provide valid identifier. (Numbers only)")
		return 0, appError
	}
	return hashId, nil
}

//UppercasePresent takes in a password and iterates over it returning whether it contains an uppercase letter
//and a number.
func UppercasePresent(password string) bool {
	for _, c := range password {
		if unicode.IsUpper(c) {
			return true
		}
	}
	return false
}
