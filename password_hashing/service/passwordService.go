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
	FindById(req dto.NewHashRequest) (*dto.NewHashResponse, *errs.AppError)
}

type DefaultPasswordService struct {
	repo domain.HashRepository
}

// NewHash takes in a NewHashRequest dto and passes the information to the domain in order to convert and save.
func (service DefaultPasswordService) NewHash(req dto.NewHashRequest) (*dto.NewHashResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	password := domain.Password{
		PasswordString: req.PasswordString,
		Id:             req.Id,
	}
	hash := domain.Hash{HashString: "", Id: req.Id}
	newPassword, err := service.repo.Save(password, hash)
	if err != nil {
		return nil, err
	}
	response := newPassword.ToNewHashResponseDto()

	return &response, nil
}

func (service DefaultPasswordService) FindById(req dto.NewHashRequest) (*dto.NewHashResponse, *errs.AppError) {
	err := req.ValidateId()
	if err != nil {
		return nil, err
	}
	targetHash, err := service.repo.FindBy(req.Id)
	if err != nil {
		return nil, err
	}
	response := targetHash.ToNewHashResponseDto()
	return &response, nil
}

func NewPasswordService(repository domain.HashRepository) DefaultPasswordService {
	return DefaultPasswordService{repository}
}
