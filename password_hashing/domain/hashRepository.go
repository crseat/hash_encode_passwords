package domain

import (
	"password_hashing/errs"
)

type HashRepositoryMap struct {
	passwords map[int64]string
}

// Save takes a new Password object and enters the id into the map. Sends the string over to be hashed by HashPassword().
func (hashRepo HashRepositoryMap) Save(password Password, hash Hash) (*Hash, *errs.AppError) {
	hashRepo.passwords[password.Id] = "Password hashing not yet complete"
	//hashRepo.HashPassword(password)
	//Don't need hashstring yet
	hash.HashString = ""
	hash.Id = password.Id
	return &hash, nil
}

func (hashRepo HashRepositoryMap) FindBy(identifier int64) (*Hash, *errs.AppError) {
	//TODO implement me
	panic("implement me")
}

func (hashRepo HashRepositoryMap) UpdateHash(identifier int64) *errs.AppError {
	//TODO implement me
	panic("implement me")
}

func (hashRepo HashRepositoryMap) HashPassword(password Password) *errs.AppError {
	//TODO: Hash Password
	err := hashRepo.UpdateHash(password.Id)
	return err
}

func NewHashRepository(passwords map[int64]string) HashRepositoryMap {
	return HashRepositoryMap{passwords: passwords}
}

func FindBy(identifier int) (string, *errs.AppError) {
	//TODO implement me
	panic("implement me")
}
