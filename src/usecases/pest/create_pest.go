package pest

import (
	"errors"
	"pesthub/contracts"
	"pesthub/entities"
)

var (
	ErrDuplicated = errors.New("this record already exists")
	ErrValidation = errors.New("invalid input")
)

type CreatePestInput struct {
	CommonName string
}

type CreatePest struct {
	save      contracts.SavePest
	getByName contracts.GetPestByName
}

func NewCreatePest(implSave contracts.SavePest, implGetByName contracts.GetPestByName) *CreatePest {
	return &CreatePest{save: implSave, getByName: implGetByName}
}

func (uc *CreatePest) Execute(userInput *CreatePestInput) (int64, error) {
	if err := uc.validateUserInput(userInput); err != nil {
		return 0, err
	}

	if err := uc.validateNameAlreadyExists(userInput.CommonName); err != nil {
		return 0, err
	}

	pest := convertToEntity(userInput)
	id, err := uc.save(pest)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (uc *CreatePest) validateNameAlreadyExists(name string) error {
	if pest, err := uc.getByName(name); err != nil {
		return err
	} else if pest != nil {
		return ErrDuplicated
	}
	return nil
}

func (*CreatePest) validateUserInput(data *CreatePestInput) error {
	if len(data.CommonName) < 3 {
		return ErrValidation
	}
	return nil
}

func convertToEntity(input *CreatePestInput) *entities.Pest {
	return &entities.Pest{
		CommonName: input.CommonName,
	}
}
