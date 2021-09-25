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

func createTestableInput() *disorder.RegisterDisorderInput {
	return &disorder.RegisterDisorderInput{
		Name: faker.Name(),
	}
}

func resetEnv() {
	env.DisorderStore = memdb.NewMemoryDisorderStore()
	env.MessageProvider = testmsgs.NewTestableMessageProvider()
}

func TestRegisterDisorder(t *testing.T) {
	t.Run("it should follow the interface", func(_ *testing.T) {
		var _ disorder.RegisterDisorderUseCase = disorder.RegisterDisorder
	})

	t.Run("it should return generated id", func(t *testing.T) {
		resetEnv()
		input := createTestableInput()
		output, err := disorder.RegisterDisorder(input)
		assert.NoError(t, err)
		assert.NotEmpty(t, output.Id)
	})

	t.Run("it should return error when name already exists", func(t *testing.T) {
		resetEnv()
		input := createTestableInput()
		entity := input.ToEntity()
		env.DisorderStore.Save(entity)

		_, err := disorder.RegisterDisorder(input)
		assert.Error(t, err)
		assert.IsType(t, failures.UseCaseError{}, err)
		assert.Equal(t, disorder.MsgNameAlreadyExists, err.Error())
	})

	t.Run("it should validate name", func(t *testing.T) {
		input := createTestableInput()
		input.Name = "" // invalid name

		_, err := disorder.RegisterDisorder(input)
		assert.Error(t, err)
		assert.IsType(t, failures.ValidationError{}, err)

	})

}
