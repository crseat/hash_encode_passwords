//Package domain handles all the business logic (how hashes are created, stored, and changed)
package domain

import (
	"password_hashing/dto"
	"password_hashing/errs"
	"sync"
	"time"
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
	return dto.NewHashResponse{
		HashString: hash.HashString,
		HashId:     hash.Id,
	}
}

//HashRepository defines the interface for saving and retrieving Password and Hash objects.
type HashRepository interface {
	Save(Password, Hash, time.Time, *sync.WaitGroup) (*Hash, *errs.AppError)
	FindBy(identifier int64) (*Hash, *errs.AppError)
	UpdateHash(int64, string, time.Time)
	HashPassword(Password) (string, *errs.AppError)
	UpdateAverage(time.Time) *errs.AppError
	IncTotal() *errs.AppError
}
