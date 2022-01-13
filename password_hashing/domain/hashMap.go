package domain

import (
	"password_hashing/errs"
)

type HashRepositoryDb struct {
	passwords map[int]string
}

// Save takes a new Password object and enters the id into the map. Sends the string over to be hashed by HashPassword().
func (hashRepo HashRepositoryDb) Save(password Password, hash Hash) (*Hash, *errs.AppError) {
	hashRepo.passwords[password.Id] = "Password hashing not yet complete"
	hashRepo.HashPassword(password)
	return password.Id, nil
}

func (hashRepo HashRepositoryDb) FindBy(identifier int) (*Hash, *errs.AppError) {
	//TODO implement me
	panic("implement me")
}

func (hashRepo HashRepositoryDb) UpdateHash(identifier int) *errs.AppError {
	//TODO implement me
	panic("implement me")
}

func (hashRepo HashRepositoryDb) HashPassword(password Password) *errs.AppError {
	//TODO: Hash Password
	err := h.UpdateHash(password.Id)
	return err
}

func NewHashRepository(passwords map[int]string) HashRepositoryDb {
	return HashRepositoryDb{passwords: passwords}
}

func FindBy(identifier int) (string, *errs.AppError) {

}
