// Package provides simple memory persitence for testing
// It's not safe for use in production
package memorydb

import (
	"math"
	"math/rand"
	"pesthub/entities"
)

type MemoryDisorderStore struct {
	disorders []*entities.Disorder
}

func NewMemoryDisorderStore() *MemoryDisorderStore {
	return &MemoryDisorderStore{
		disorders: make([]*entities.Disorder, 0, 10),
	}
}

func (s *MemoryDisorderStore) Save(disorder *entities.Disorder) (*entities.Disorder, error) {
	disorder.Code = randomize()
	s.disorders = append(s.disorders, clone(*disorder))
	return disorder, nil
}

func (s *MemoryDisorderStore) ExistsName(name string) (bool, error) {
	for _, disorder := range s.disorders {
		if name == disorder.Name {
			return true, nil
		}
	}
	return false, nil
}

func randomize() uint64 {
	number := rand.Int63n(math.MaxInt64)
	return uint64(number)
}

func clone(item entities.Disorder) *entities.Disorder {
	// will always work ?
	return &item
}
