package pest

import (
	"errors"
	"pesthub/entities"
)

type CheckAlreadyExistsInput struct {
	ScientificName string
	CommonName     string
}

type CheckAlreadyExists func(data *CheckAlreadyExistsInput) error

func NewCheckAlreadyExists(
	findAllHavingNames FindAllHavingNames,
) CheckAlreadyExists {
	return func(data *CheckAlreadyExistsInput) error {
		pests, err := findAllHavingNames(data.CommonName)
		if err != nil {
			return err
		}

		if len(pests) > 0 {
			return errors.New("pest already exists")
		}

		return nil
	}
}

type FindOneHavingName func(name string) (*entities.Pest, error)
type FindAllHavingNames func(names ...string) ([]*entities.Pest, error)
