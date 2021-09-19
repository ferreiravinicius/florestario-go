package contracts

import "pesthub/entities"

type DisorderStore interface {
	Save(*entities.Disorder) (*entities.Disorder, error)
	ExistsName(name string) (bool, error)
	FindAll() ([]entities.Disorder, error)
}
