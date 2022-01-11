//Package domain handles all the business logic (how passwords are created, stored, and changed)
package domain

import (
	"password_hashing/dto"
	"password_hashing/errs"
)

type Password struct {
	Id   int
	Hash string
}

//ToNewHashResponseDto takes a Password object and converts it into an appropriate response to the client.
func (p Password) ToNewHashResponseDto() dto.NewHashResponse {
	return dto.NewHashResponse{Hash: p.Hash}
}

//HashRepository defines the interface for saving and retrieving Password objects.
type HashRepository interface {
	Save(Password) (*Password, *errs.AppError)
	FindBy(identifier int) (*Password, *errs.AppError)
}
