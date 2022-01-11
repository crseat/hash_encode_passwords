package domain

import (
	"password_hashing/dto"
	"password_hashing/errs"
)

type Password struct {
	Id   int
	Hash string
}

func (p Password) ToNewHashResponseDto() dto.NewHashResponse {
	return dto.NewHashResponse{Hash: p.Hash}
}

type HashRepository interface {
	Save(Password) (*Password, *errs.AppError)
	FindBy(identifier int) (*Password, *errs.AppError)
}
