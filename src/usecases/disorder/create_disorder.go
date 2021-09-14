package disorder

import (
	"pesthub/commons/errors"
	"pesthub/contracts/store"
	"pesthub/entities"
)

type CreateDisorderInput struct {
	Name string
}

type ICreateDisorder func(data *CreateDisorderInput) (int64, error)

func NewCreateDisorder(
	saveDisorder store.SaveDisorder,
	existsDisorderByName store.ExistsDisorderByName,
) ICreateDisorder {
	return func(data *CreateDisorderInput) (int64, error) {
		if err := validate(data); err != nil {
			return 0, err
		}

		existsName, err := existsDisorderByName(data.Name)
		if err != nil {
			return 0, errors.Unexpected(err)
		}
		if existsName {
			return 0, errors.Business("disorder.already.exists.name")
		}

		disorder := convert(data)
		id, err := saveDisorder(disorder)
		if err != nil {
			return 0, errors.Unexpected(err)
		}

		return id, nil
	}
}

func convert(userInput *CreateDisorderInput) *entities.Disorder {
	return &entities.Disorder{}
}

// maybe decouple this ?
func validate(data *CreateDisorderInput) error {
	if len(data.Name) == 0 {
		return errors.Business("field.required.name")
	}
	return nil
}
