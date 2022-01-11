package service

import (
	"password_hashing/domain"
	"password_hashing/dto"
	"password_hashing/errs"
)

type PasswordService interface {
	NewPassword(request dto.NewHashRequest) (*dto.NewHashResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.HashRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewHashRequest) (*dto.NewHashResponse, *errs.AppError) {
	return
}
