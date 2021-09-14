package disorder

import (
	"errors"
	"pesthub/contracts/store"
	"pesthub/entities"
)

type CreateDisorderInput struct {
	BionomialName string
	Name          string
	Description   string
	Group         string //TODO: use entity ?
}

type ICreateDisorder func(data *CreateDisorderInput) (int64, error)

func NewCreateDisorder(
	save store.SaveDisorder,
	checkAlreadyExists ICheckAlreadyExists,
) ICreateDisorder {
	return func(data *CreateDisorderInput) (int64, error) {
		if err := validate(data); err != nil {
			return 0, err
		}

		if err := checkAlreadyExists(&CheckAlreadyExistsInput{Name: data.Name}); err != nil {
			return 0, err
		}

		disorder := convert(data)
		id, err := save(disorder)
		if err != nil {
			return 0, err
		}

		return id, nil
	}
}

func convert(userInput *CreateDisorderInput) *entities.Disorder {
	return &entities.Disorder{}
}

// maybe change decouple this ?
func validate(data *CreateDisorderInput) error {
	if len(data.Name) == 0 {
		return errors.New("name is required")
	}
	if len(data.BionomialName) == 0 {
		return errors.New("binomial name is required")
	}
	return nil
}
