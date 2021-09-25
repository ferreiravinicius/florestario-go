package disorder_test

import (
	"pesthub/adapters/memdb"
	"pesthub/adapters/testmsgs"
	"pesthub/env"
	"pesthub/failures"
	"pesthub/usecases/disorder"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func makeTestableRegisterDisorderInput() *disorder.RegisterDisorderInput {
	return &disorder.RegisterDisorderInput{
		Name: faker.Name(),
	}
}

func configureEnv() {
	env.DisorderStore = memdb.NewMemoryDisorderStore()
	env.MessageProvider = testmsgs.NewTestableMessageProvider()
}

func TestRegisterDisorder(t *testing.T) {
	t.Run("it should follow the interface", func(_ *testing.T) {
		var _ disorder.RegisterDisorderUseCase = disorder.RegisterDisorder
	})

	t.Run("it should return generated id", func(t *testing.T) {
		configureEnv()
		input := makeTestableRegisterDisorderInput()
		output, err := disorder.RegisterDisorder(input)
		assert.NoError(t, err)
		assert.NotEmpty(t, output.Id)
	})

	t.Run("it should return error when name already exists", func(t *testing.T) {
		configureEnv()
		input := makeTestableRegisterDisorderInput()
		entity := input.ToEntity()
		env.DisorderStore.Save(entity)

		_, err := disorder.RegisterDisorder(input)
		assert.Error(t, err)
		assert.IsType(t, failures.UseCaseError{}, err)
		assert.Equal(t, disorder.MsgNameAlreadyExists, err.Error())
	})
}
