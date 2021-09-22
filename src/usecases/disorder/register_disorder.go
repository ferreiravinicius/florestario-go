package disorder

import (
	"pesthub/contracts"
	"pesthub/entities"
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

type RegisterDisorderUseCase interface {
	Execute(input *RegisterDisorderInput) (*RegisterDisorderOutput, error)
}

type RegisterDisorder struct {
	store    contracts.DisorderStore
	messages contracts.Messages
}

func NewRegisterDisorder(store contracts.DisorderStore, messages contracts.Messages) RegisterDisorderUseCase {
	return &RegisterDisorder{
		store,
		messages,
	}
}

func (usecase *RegisterDisorder) Execute(disorderInput *RegisterDisorderInput) (*RegisterDisorderOutput, error) {
	exists, err := usecase.store.ExistsName(disorderInput.Name)
	if err != nil {
		return nil, failures.Internal(err)
	}
	if exists {
		message := usecase.messages.GetText(MsgNameAlreadyExists)
		return nil, failures.UseCase(message)
	}

	disorder := disorderInput.ToEntity()
	disorder, err = usecase.store.Save(disorder)
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
