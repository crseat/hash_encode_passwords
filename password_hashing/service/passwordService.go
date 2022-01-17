//Package service defines and implements the actions available to be performed by the user.
package service

import (
	"password_hashing/domain"
	"password_hashing/dto"
	"password_hashing/errs"
	"sync"
	"time"
)

//PasswordService processes requests for new and existing hashes.
type PasswordService interface {
	NewHash(dto.NewHashRequest, time.Time, *sync.WaitGroup) (*dto.NewHashResponse, *errs.AppError)
	FindById(req dto.NewHashRequest) (*dto.NewHashResponse, *errs.AppError)
	IncTotal() *errs.AppError
	UpdateAverage(startTime time.Time) *errs.AppError
}

type DefaultPasswordService struct {
	Repo domain.HashRepository
}

// NewHash takes in a NewHashRequest dto and passes the information to the domain in order to convert and save.
func (service DefaultPasswordService) NewHash(req dto.NewHashRequest, startTime time.Time, wg *sync.WaitGroup) (*dto.NewHashResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	password := domain.Password{
		PasswordString: req.PasswordString,
		Id:             req.Id,
	}
	hash := domain.Hash{HashString: "", Id: req.Id}

	//Call service to save the new password in our repo
	newPassword, err := service.Repo.Save(password, hash, startTime, wg)
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
	targetHash, err := service.Repo.FindBy(req.Id)
	if err != nil {
		return nil, err
	}
	response := targetHash.ToNewHashResponseDto()
	return &response, nil
}

func (service DefaultPasswordService) IncTotal() *errs.AppError {
	err := service.Repo.IncTotal()
	if err != nil {
		return err
	}
	return nil
}

func (service DefaultPasswordService) UpdateAverage(startTime time.Time) *errs.AppError {
	err := service.Repo.UpdateAverage(startTime)
	if err != nil {
		return err
	}
	return nil
}

func NewPasswordService(hashRepository domain.HashRepository) DefaultPasswordService {
	return DefaultPasswordService{hashRepository}
}
