package disorder

import (
	"pesthub/entities"
	"pesthub/env"
	"pesthub/failures"
	"strconv"
)

const (
	MsgNameAlreadyExists = "disorder.name.exists"
)

type RegisterDisorderInput struct {
	Name string
}

type RegisterDisorderOutput struct {
	Code string
}

func RegisterDisorder(data *RegisterDisorderInput) (*RegisterDisorderOutput, error) {
	exists, err := env.DisorderStore.ExistsName(data.Name)
	if err != nil {
		return nil, failures.Internal(err)
	}
	if exists {
		message := env.Messages.GetText(MsgNameAlreadyExists)
		return nil, failures.UseCase(message)
	}

	disorder := data.ToEntity()
	disorder, err = env.DisorderStore.Save(disorder)
	if err != nil {
		return nil, failures.Internal(err)
	}

	return &RegisterDisorderOutput{
		Code: strconv.FormatUint(disorder.Code, 10),
	}, nil
}

func (in *RegisterDisorderInput) ToEntity() *entities.Disorder {
	return &entities.Disorder{Name: in.Name}
}
