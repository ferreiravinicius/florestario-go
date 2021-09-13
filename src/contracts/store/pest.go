package store

import "pesthub/entities"

type SaveDisorder func(pest *entities.Disorder) (int64, error)
type FindDisorderHavingNames func(names ...string) ([]*entities.Disorder, error)

// type FindPestHavingName func(name string) (*entities.Pest, error)
