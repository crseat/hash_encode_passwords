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

// FindById takes in a NewHashRequest, queries the repo for the hash using the corresponding id, and then converts
// response into NewHashResponse
func (service DefaultPasswordService) FindById(req dto.NewHashRequest) (*dto.NewHashResponse, *errs.AppError) {
	targetHash, err := service.Repo.FindBy(req.Id)
	if err != nil {
		return nil, err
	}
	response := targetHash.ToNewHashResponseDto()
	return &response, nil
}

// IncTotal calls the service to increment the counter for total amount of Post Requests
func (service DefaultPasswordService) IncTotal() *errs.AppError {
	err := service.Repo.IncTotal()
	if err != nil {
		return err
	}
	return nil
}

// UpdateAverage takes in the start time of the POST request and calls the service to update average time taken to
// process POST requests
func (service DefaultPasswordService) UpdateAverage(startTime time.Time) *errs.AppError {
	err := service.Repo.UpdateAverage(startTime)
	if err != nil {
		return err
	}
	return nil
}

// NewPasswordService creates new DefaultPasswordService using the passed in hashRepo
func NewPasswordService(hashRepository domain.HashRepository) DefaultPasswordService {
	return DefaultPasswordService{hashRepository}
}
