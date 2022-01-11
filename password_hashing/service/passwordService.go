//Package service defines and implements the actions available to be performed by the user.
package service

import (
	"password_hashing/domain"
	"password_hashing/dto"
	"password_hashing/errs"
)

//PasswordService processes requests for new and existing hashes.
type PasswordService interface {
	NewHash(request dto.NewHashRequest) (*dto.NewHashResponse, *errs.AppError)
}

type DefaultPasswordService struct {
	repo domain.HashRepository
}

//func (s DefaultPasswordService) NewHash(req dto.NewHashRequest) (*dto.NewHashResponse, *errs.AppError) {
//	return
//}
