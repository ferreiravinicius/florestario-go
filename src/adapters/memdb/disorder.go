// Package provides simple memory persitence for testing
// Not safe for use in production
package memdb

import (
	"math"
	"math/rand"
	"pesthub/contracts"
	"pesthub/entities"
)

type memoryDisorderStore struct {
	disorders []entities.Disorder
}

func NewMemoryDisorderStore() contracts.DisorderStore {
	return &memoryDisorderStore{
		disorders: make([]entities.Disorder, 0, 10),
	}
}

func (s *memoryDisorderStore) Save(disorder *entities.Disorder) error {
	disorder.Id = randomize()
	s.disorders = append(s.disorders, *disorder)
	return nil
}

func (s *memoryDisorderStore) ExistsName(name string) (bool, error) {
	for _, disorder := range s.disorders {
		if name == disorder.Name {
			return true, nil
		}
	}
	return false, nil
}

func (s *memoryDisorderStore) FindAll() ([]entities.Disorder, error) {
	return s.disorders, nil
}

func randomize() uint64 {
	number := rand.Int63n(math.MaxInt64)
	return uint64(number)
}
