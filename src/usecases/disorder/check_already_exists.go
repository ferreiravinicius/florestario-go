package disorder

import (
	"errors"
	"pesthub/contracts/store"
)

type CheckAlreadyExistsInput struct {
	Names   []string
	Exclude string //todo: exclude id
}

type ICheckAlreadyExists func(data *CheckAlreadyExistsInput) error

func NewCheckAlreadyExists(
	findDisorderHavingNames store.FindDisorderHavingNames,
) ICheckAlreadyExists {
	return func(data *CheckAlreadyExistsInput) error {
		disorder, err := findDisorderHavingNames(data.Names...)
		if err != nil {
			return err
		}

		if len(disorder) > 0 {
			return errors.New("disorder already exists")
		}

		return nil
	}
}
