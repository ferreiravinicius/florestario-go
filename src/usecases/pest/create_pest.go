package pest

import (
	"errors"
	"pesthub/contracts/store"
	"pesthub/entities"
)

type CreatePestInput struct {
	BionomialName string
	Name          string
}

type ICreatePest func(data *CreatePestInput) (int64, error)

func NewCreatePest(
	save store.SavePest,
	checkAlreadyExists ICheckAlreadyExists,
) ICreatePest {
	return func(data *CreatePestInput) (int64, error) {
		if err := validate(data); err != nil {
			return 0, err
		}

		if err := checkAlreadyExists(&CheckAlreadyExistsInput{
			ScientificName: "",
			CommonName:     "",
		}); err != nil {
			return 0, err
		}

		pest := convert(data)
		id, err := save(pest)
		if err != nil {
			return 0, err
		}

		return id, nil
	}
}

func convert(userInput *CreatePestInput) *entities.Pest {
	return &entities.Pest{}
}

// maybe change decouple this ?
func validate(data *CreatePestInput) error {
	if len(data.Name) == 0 {
		return errors.New("name is required")
	}
	if len(data.BionomialName) == 0 {
		return errors.New("binomial name is required")
	}
	return nil
}
