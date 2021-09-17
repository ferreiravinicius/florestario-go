package disorder_test

import (
	"math"
	"math/rand"
	"pesthub/deps"
	"pesthub/entities"
)

type InMemoryDisorderStore struct {
	disorders []*entities.Disorder
}

func (s *InMemoryDisorderStore) Save(disorder *entities.Disorder) (*entities.Disorder, error)  {
	disorder.Code = MakeRandomCode()
	s.disorders = append(s.disorders, disorder)
	return disorder, nil
}

func MakeRandomCode() uint64 {
	return rand.Int63n(math.MaxInt64)
}

Save(*entities.Disorder) (*entities.Disorder, error)
	ExistsName(name string) (bool, error

func init() {
	deps.DisorderStore = 
}