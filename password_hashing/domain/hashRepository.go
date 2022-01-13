package domain

import (
	"password_hashing/errs"
)

type HashRepository struct {
	passwords map[int64]string
}

// Save takes a new Password object and enters the id into the map. Sends the string over to be hashed by HashPassword().
func (hashRepo HashRepository) Save(password Password, hash Hash) (*Hash, *errs.AppError) {
	hashRepo.passwords[password.Id] = "Password hashing not yet complete"
	hashRepo.HashPassword(password)
	return password.Id, nil
}

func (hashRepo HashRepository) FindBy(identifier int64) (*Hash, *errs.AppError) {
	//TODO implement me
	panic("implement me")
}

func (hashRepo HashRepository) UpdateHash(identifier int64) *errs.AppError {
	//TODO implement me
	panic("implement me")
}

func (hashRepo HashRepository) HashPassword(password Password) *errs.AppError {
	//TODO: Hash Password
	err := hashRepo.UpdateHash(password.Id)
	return err
}

func NewHashRepository(passwords map[int]string) HashRepository {
	return HashRepository{passwords: passwords}
}

func FindBy(identifier int) (string, *errs.AppError) {

}
