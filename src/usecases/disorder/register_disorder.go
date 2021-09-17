package disorder

import (
	"errors"
	"pesthub/deps"
	"pesthub/entities"
)

const (
	MsgNameAlreadyExists = "disorder.name.exists"
)

type RegisterDisorderInput struct {
	Name string
}

//TODO: change output to PresentableDisorder
func RegisterDisorder(data *RegisterDisorderInput) (*entities.Disorder, error) {
	exists, err := deps.DisorderStore.ExistsName(data.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		message := deps.Messages.GetText(MsgNameAlreadyExists)
		return nil, errors.New(message)
	}

	disorder := makeDisorder(data)
	disorder, err = deps.DisorderStore.Save(disorder)
	if err != nil {
		return nil, err
	}

	return disorder, nil
}

func makeDisorder(data *RegisterDisorderInput) *entities.Disorder {
	return &entities.Disorder{
		Name: data.Name,
	}
}
