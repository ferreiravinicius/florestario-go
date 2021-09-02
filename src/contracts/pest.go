package contracts

import "pesthub/entities"

type SavePest func(pest *entities.Pest) (int64, error)
type GetPestByName func(name string) (*entities.Pest, error)
