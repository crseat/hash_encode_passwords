package domain

import (
	"crypto/sha512"
	"encoding/base64"
	"password_hashing/errs"
	"time"
)

type HashRepositoryMap struct {
	hashes map[int64]string
}

// Save takes a new Password object and enters the id into the map. Sends the string over to be hashed by HashPassword().
func (hashRepo HashRepositoryMap) Save(password Password, hash Hash) (*Hash, *errs.AppError) {

	//set a temporary value for the hashString incase the user tries to query before 5 seconds
	hashRepo.hashes[password.Id] = "Password hashing not yet complete"

	// We only need the id to be returned to the when first saving new password
	hash.HashString = ""
	hash.Id = password.Id

	// Hash password after 5 seconds
	time.AfterFunc(5*time.Second, func() {
		hashString, err := hashRepo.HashPassword(password)
		if err != nil {
			hashString = "There was an error while hashing password"
		}
		// Update the map with the calculated hash.
		hashRepo.UpdateHash(password.Id, hashString)
	})

	return &hash, nil
}

func (hashRepo HashRepositoryMap) FindBy(identifier int64) (*Hash, *errs.AppError) {
	hash := Hash{
		HashString: hashRepo.hashes[identifier],
		Id:         identifier,
	}
	return &hash, nil
}

func (hashRepo HashRepositoryMap) UpdateHash(identifier int64, hashString string) {
	hashRepo.hashes[identifier] = hashString
}

func (hashRepo HashRepositoryMap) HashPassword(password Password) (string, *errs.AppError) {
	hash := sha512.Sum512([]byte(password.PasswordString))
	hashEncoded := "{SHA512}" + base64.StdEncoding.EncodeToString(hash[:])
	return hashEncoded, nil
}

func NewHashRepository(passwords map[int64]string) HashRepositoryMap {
	return HashRepositoryMap{hashes: passwords}
}

func FindBy(identifier int) (string, *errs.AppError) {
	//TODO implement me
	panic("implement me")
}
