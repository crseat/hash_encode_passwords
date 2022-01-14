//Package domain handles all the business logic (how passwords are created, stored, and changed)
package domain

import (
	"password_hashing/dto"
	"password_hashing/errs"
)

type Password struct {
	PasswordString string
	Id             int64
}

type Hash struct {
	HashString string
	Id         int64
}

//ToNewHashResponseDto takes a Hash object and converts it into an appropriate response to the client.
func (hash Hash) ToNewHashResponseDto() dto.NewHashResponse {
	return dto.NewHashResponse{HashId: hash.Id}
}

//HashRepository defines the interface for saving and retrieving Password and Hash objects.
type HashRepository interface {
	Save(Password, Hash) (*Hash, *errs.AppError)
	FindBy(identifier int64) (*Hash, *errs.AppError)
	UpdateHash(identifier int64) *errs.AppError
	HashPassword(password Password) *errs.AppError
}
