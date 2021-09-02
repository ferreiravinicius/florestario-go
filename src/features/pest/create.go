package pest

import (
	"errors"
	"pesthub/entities"
)

type InsertCommand func(pest *entities.Pest) (int64, error)

type CreateDeps struct {
	InsertCommand
}
type CreateInput struct {
	CommonName string
}

func (i *CreateInput) ToEntity() *entities.Pest {
	return &entities.Pest{
		CommonName: i.CommonName,
	}
}

func Create(deps CreateDeps, data *CreateInput) (int64, error) {
	if err := validate(data); err != nil {
		return 0, err
	}
	pest := data.ToEntity()
	id, err := deps.InsertCommand(pest)
	if err != nil {
		return 0, err
	}
	return id, nil
}

var ErrValidation = errors.New("invalid input")

func validate(data *CreateInput) error {
	if len(data.CommonName) < 3 {
		return ErrValidation
	}
	return nil
}
