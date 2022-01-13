//Package domain handles all the business logic (how passwords are created, stored, and changed)
package domain

import (
	"password_hashing/dto"
	"password_hashing/errs"
)

type Password struct {
	PasswordString string
	Id             int
}

type Hash struct {
	HashString string
	Id         int
}

//ToNewHashResponseDto takes a Hash object and converts it into an appropriate response to the client.
func (hash Hash) ToNewHashResponseDto() dto.NewHashResponse {
	return dto.NewHashResponse{Hash: hash.HashString}
}

//HashRepository defines the interface for saving and retrieving Password and Hash objects.
type HashRepository interface {
	Save(Password) (*Hash, *errs.AppError)
	FindBy(identifier int) (*Hash, *errs.AppError)
	UpdateHash(identifier int) *errs.AppError
	HashPassword(password Password) *errs.AppError
}
