package disorder

import (
	"errors"
	"pesthub/entities"
	"pesthub/scope"
)

type RegisterDisorderInput struct {
	Name string
}

func RegisterDisorder(data *RegisterDisorderInput) (*entities.Disorder, error) {
	exists, err := scope.DisorderStore.ExistsName(data.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		message := scope.Messages.GetText("code")
		return nil, errors.New(message)
	}

	disorder := makeDisorder(data)
	disorder, _ = scope.DisorderStore.Save(disorder)
	return disorder, nil
}

func makeDisorder(data *RegisterDisorderInput) *entities.Disorder {
	return &entities.Disorder{
		Name: data.Name,
	}
}
