package domain

import (
	"sync"
	"testing"
	"time"
)

func Test_should_assert_proper_saving_of_new_Password(t *testing.T) {
	password := Password{
		PasswordString: "angryMonkey",
		Id:             1,
	}
	hash := Hash{
		HashString: "",
		Id:         1,
	}
	startTime := time.Now()
	wg := &sync.WaitGroup{}

	hashMap := NewHashRepository()

	hashMap.IncTotal()

	hashMap.Save(password, hash, startTime, wg)
	test := hashMap.Hashes[1]
	_ = test
	if hashMap.Hashes[1] != "Password hashing not yet complete" {
		t.Error("Invalid message while testing password hashing")
	}
	//Need to wait for hashing to complete
	time.Sleep(6 * time.Second)
	if hashMap.Hashes[1] != "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==" {
		t.Error("Invalid message while testing password hashing")
	}
}
