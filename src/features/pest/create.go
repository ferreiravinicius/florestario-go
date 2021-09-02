package pest

import (
	"errors"
	"pesthub/contracts"
	"pesthub/entities"
)

type CreatePestInput struct {
	CommonName string
}

type CreatePest struct {
	insert contracts.InsertPest
}

func NewCreatePest(implInsert contracts.InsertPest) *CreatePest {
	return &CreatePest{insert: implInsert}
}

func (o *CreatePest) Execute(pestInput *CreatePestInput) (int64, error) {
	if err := validate(pestInput); err != nil {
		return 0, err
	}
	pest := convert(pestInput)
	id, err := o.insert(pest)
	if err != nil {
		return 0, err
	}
	return id, nil
}

var ErrValidation = errors.New("invalid input")

func validate(data *CreatePestInput) error {
	if len(data.CommonName) < 3 {
		return ErrValidation
	}
	return nil
}

func convert(input *CreatePestInput) *entities.Pest {
	return &entities.Pest{
		CommonName: input.CommonName,
	}
}
