package disorder

import (
	"pesthub/entities"
	"pesthub/env"
	"pesthub/failures"
	"pesthub/validators"
	"strconv"
)

const (
	MsgNameAlreadyExists = "disorder.name.exists"
)

type RegisterDisorderInput struct {
	Name string `json:"name"`
}

type RegisterDisorderOutput struct {
	Id string `json:"id"`
}

type RegisterDisorderUseCase func(input *RegisterDisorderInput) (*RegisterDisorderOutput, error)

func RegisterDisorder(disorderInput *RegisterDisorderInput) (*RegisterDisorderOutput, error) {

	if err := validators.Name(disorderInput.Name); err != nil {
		return nil, err
	}

	exists, err := env.DisorderStore.ExistsName(disorderInput.Name)
	if err != nil {
		return nil, failures.Internal(err)
	}
	if exists {
		message := env.MessageProvider.Get(MsgNameAlreadyExists)
		return nil, failures.UseCase(message)
	}

	entity := disorderInput.ToEntity()
	if err = env.DisorderStore.Save(entity); err != nil {
		return nil, failures.Internal(err)
	}

	return &RegisterDisorderOutput{
		Id: strconv.FormatUint(entity.Id, 10),
	}, nil
}

func (in *RegisterDisorderInput) ToEntity() *entities.Disorder {
	return &entities.Disorder{Name: in.Name}
}
