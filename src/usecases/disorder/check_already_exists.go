package disorder

import (
	"pesthub/commons/errors"
	"pesthub/contracts/store"
)

type CheckAlreadyExistsInput struct {
	Name    string
	Exclude string //TODO: exclude id
}

type ICheckAlreadyExists func(data *CheckAlreadyExistsInput) error

func NewCheckAlreadyExists(
	getDisorderByName store.GetDisorderByName,
) ICheckAlreadyExists {
	return func(data *CheckAlreadyExistsInput) error {
		disorder, err := getDisorderByName(data.Name)
		if err != nil {
			return errors.Unexpected(err)
		}
		if disorder != nil {
			return errors.Business("disorder.exists.name")
		}
		return nil
	}
}
