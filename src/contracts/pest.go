package contracts

import "pesthub/entities"

type InsertPest func(pest *entities.Pest) (int64, error)
