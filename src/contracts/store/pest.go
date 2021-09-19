package store

import "pesthub/entities"

type SaveDisorder func(pest *entities.Disorder) (int64, error)
type GetDisorderByName func(name string) (*entities.Disorder, error)
type ExistsDisorderByName func(name string) (bool, error)

type DisorderStore interface {
	Save(*entities.Disorder) (*entities.Disorder, error)
	ExistsName(name string) (bool, error)
	FindAll() ([]entities.Disorder, error)
}
