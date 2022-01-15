//Package domain handles all the business logic (how hashes are created, stored, and changed)
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

type Stats struct {
	Total   int64
	Average int64
}

//ToNewHashResponseDto takes a Hash object and converts it into an appropriate response to the client.
func (hash Hash) ToNewHashResponseDto() dto.NewHashResponse {
	return dto.NewHashResponse{
		HashString: hash.HashString,
		HashId:     hash.Id,
	}
}

//ToNewStatsResponseDto takes a Stats object and converts it into an appropriate response to the client.
func (stats Stats) ToNewStatsResponseDto() dto.NewStatsResponse {
	return dto.NewStatsResponse{
		Total:   stats.Total,
		Average: stats.Average,
	}
}

//HashRepository defines the interface for saving and retrieving Password and Hash objects.
type HashRepository interface {
	Save(Password, Hash) (*Hash, *errs.AppError)
	FindBy(identifier int64) (*Hash, *errs.AppError)
	UpdateHash(int64, string)
	HashPassword(password Password) (string, *errs.AppError)
	GetStats() (*Stats, *errs.AppError)
}
