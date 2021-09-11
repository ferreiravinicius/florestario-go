package pest

import (
	"pesthub/contracts/store"
	"pesthub/entities"
)

type CreatePestInput struct {
}

type ICreatePest func(data *CreatePestInput) (int64, error)

func NewCreatePest(
	save store.SavePest,
	checkAlreadyExists CheckAlreadyExists,
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

func validate(userInput *CreatePestInput) error {
	return nil
}
