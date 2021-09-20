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

func init() {
	env.DisorderStore = memdb.NewMemoryDisorderStore()
	env.Messages = testmsgs.NewTestableMessages()
}

func makeTestableRegisterDisorderInput() *disorder.RegisterDisorderInput {
	return &disorder.RegisterDisorderInput{
		Name: faker.Name(),
	}
}

func TestRegisterDisorder(t *testing.T) {
	t.Run("it should return generated code", func(t *testing.T) {
		input := makeTestableRegisterDisorderInput()
		output, err := disorder.RegisterDisorder(input)
		assert.NoError(t, err)
		assert.NotEmpty(t, output.Code)
	})

	t.Run("it should return error when name already exists", func(t *testing.T) {
		input := makeTestableRegisterDisorderInput()
		env.DisorderStore.Save(input.ToEntity()) // save before
		_, err := disorder.RegisterDisorder(input)
		assert.Error(t, err)
		assert.IsType(t, failures.UseCaseError{}, err)
		assert.Equal(t, disorder.MsgNameAlreadyExists, err.Error())
	})
}
