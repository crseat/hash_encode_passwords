package domain

import (
	"crypto/sha512"
	"encoding/base64"
	"password_hashing/errs"
	"sync/atomic"
	"time"
)

type HashRepositoryMap struct {
	hashes map[int64]string
}

// Define and keep track of the password ids
//source: https://stackoverflow.com/questions/27917750/how-to-define-a-global-counter-in-golang-http-server
var id int64 = 0

// increments the number of the id and returns the new value
func incId() int64 {
	return atomic.AddInt64(&id, 1)
}

// returns the current value
func getId() int64 {
	return atomic.LoadInt64(&id)
}

// Save takes a new Password object and enters the id into the map. Sends the string over to be hashed by HashPassword().
func (hashRepo HashRepositoryMap) Save(password Password, hash Hash) (*Hash, *errs.AppError) {
	incId()
	passwordId := getId()
	//set a temporary value for the hashString incase the user tries to query before 5 seconds
	hashRepo.hashes[passwordId] = "Password hashing not yet complete"

	// We only need the id to be returned to the when first saving new password
	hash.HashString = ""
	hash.Id = passwordId

	// Hash password after 5 seconds
	time.AfterFunc(5*time.Second, func() {
		hashString, err := hashRepo.HashPassword(password)
		if err != nil {
			hashString = "There was an error while hashing password"
		}
		// Update the map with the calculated hash.
		hashRepo.UpdateHash(passwordId, hashString)
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
	hashEncoded := base64.StdEncoding.EncodeToString(hash[:])
	return hashEncoded, nil
}

func NewHashRepository(passwords map[int64]string) HashRepositoryMap {
	return HashRepositoryMap{hashes: passwords}
}

func (hashRepo HashRepositoryMap) GetStats() (*Stats, *errs.AppError) {
	stats := Stats{
		Total:   getId(),
		Average: 0,
	}
	return &stats, nil
}
