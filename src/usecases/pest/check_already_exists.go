package pest

import (
	"errors"
	"pesthub/contracts/store"
)

type CheckAlreadyExistsInput struct {
	ScientificName string
	CommonName     string
}

type CheckAlreadyExists func(data *CheckAlreadyExistsInput) error

func NewCheckAlreadyExists(
	findPestsHavingNames store.FindPestsHavingNames,
) CheckAlreadyExists {
	return func(data *CheckAlreadyExistsInput) error {
		pests, err := findPestsHavingNames(data.CommonName)
		if err != nil {
			return err
		}

		if len(pests) > 0 {
			return errors.New("pest already exists")
		}

		return nil
	}
}
