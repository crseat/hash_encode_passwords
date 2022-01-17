package domain

import (
	"crypto/sha512"
	"encoding/base64"
	"password_hashing/errs"
	"password_hashing/logger"
	"sync"
	"sync/atomic"
	"time"
)

type HashRepositoryMap struct {
	Hashes  map[int64]string
	Total   *int64
	Average *int64
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
func (hashRepo HashRepositoryMap) Save(password Password, hash Hash, startTime time.Time, wg *sync.WaitGroup) (*Hash, *errs.AppError) {
	incId()
	passwordId := getId()
	//set a temporary value for the hashString incase the user tries to query before 5 seconds
	hashRepo.Hashes[passwordId] = "Password hashing not yet complete"

	// We only need the id to be returned to the when first saving new password
	hash.HashString = ""
	hash.Id = passwordId

	// Hash password after 5 seconds
	wg.Add(1)
	time.AfterFunc(5*time.Second, func() {
		// Ensures graceful shutdown
		defer wg.Done()
		hashString, err := hashRepo.HashPassword(password)
		if err != nil {
			logger.ErrorLogger.Println("Hashing error: ", err)
			hashString = "There was an error while hashing password"
		}
		// Update the map with the calculated hash.
		hashRepo.UpdateHash(passwordId, hashString, startTime)
	})

	return &hash, nil
}

func (hashRepo HashRepositoryMap) FindBy(identifier int64) (*Hash, *errs.AppError) {
	hash := Hash{
		HashString: hashRepo.Hashes[identifier],
		Id:         identifier,
	}
	return &hash, nil
}

func (hashRepo HashRepositoryMap) UpdateHash(identifier int64, hashString string, startTime time.Time) {
	hashRepo.Hashes[identifier] = hashString
	hashRepo.UpdateAverage(startTime)
}

func (hashRepo HashRepositoryMap) HashPassword(password Password) (string, *errs.AppError) {
	hash := sha512.Sum512([]byte(password.PasswordString))
	hashEncoded := base64.StdEncoding.EncodeToString(hash[:])
	return hashEncoded, nil
}

func (hashRepo HashRepositoryMap) GetStats() (*Stats, *errs.AppError) {
	stats := Stats{
		Total:   *hashRepo.Total,
		Average: *hashRepo.Average,
	}
	return &stats, nil
}
func (hashRepo HashRepositoryMap) IncTotal() *errs.AppError {
	*hashRepo.Total += 1
	return nil
}

func (hashRepo HashRepositoryMap) UpdateAverage(startTime time.Time) *errs.AppError {
	duration := time.Now().Sub(startTime)
	//to calculate the new average after then nth number, you multiply the old average by n−1, add the new number,
	//and divide the total by n.
	newAverage := ((*hashRepo.Average * (*hashRepo.Total - 1)) + duration.Microseconds()) / *hashRepo.Total
	*hashRepo.Average = newAverage
	return nil
}

func NewHashRepository() HashRepositoryMap {
	//return HashRepositoryMap{Hashes: passwords}
	return HashRepositoryMap{
		Hashes:  make(map[int64]string),
		Total:   new(int64),
		Average: new(int64),
	}
}
