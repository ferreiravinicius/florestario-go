package disorder_test

import (
	"pesthub/adapters/memdb"
	"pesthub/adapters/testmsgs"
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

func TestRegisterDisorder(t *testing.T) {
	t.Run("it should return generated code", func(t *testing.T) {
		store := memdb.NewMemoryDisorderStore()
		messages := testmsgs.NewTestableMessages()
		sut := disorder.NewRegisterDisorder(store, messages)

		input := makeTestableRegisterDisorderInput()
		output, err := sut.Execute(input)
		assert.NoError(t, err)
		assert.NotEmpty(t, output.Code)
	})

	t.Run("it should return error when name already exists", func(t *testing.T) {
		store := memdb.NewMemoryDisorderStore()
		messages := testmsgs.NewTestableMessages()
		sut := disorder.NewRegisterDisorder(store, messages)

		input := makeTestableRegisterDisorderInput()
		store.Save(input.ToEntity()) // save before

		_, err := sut.Execute(input)
		assert.Error(t, err)
		assert.IsType(t, failures.UseCaseError{}, err)
		assert.Equal(t, disorder.MsgNameAlreadyExists, err.Error())
	})
}

var outscope *disorder.RegisterDisorderOutput

func BenchmarkRegisterDisorder(b *testing.B) {
	var r *disorder.RegisterDisorderOutput
	sut := disorder.NewRegisterDisorder(memdb.NewMemoryDisorderStore(), testmsgs.NewTestableMessages())
	for i := 0; i < b.N; i++ {
		input := makeTestableRegisterDisorderInput()
		r, _ = sut.Execute(input)
	}
	outscope = r
}
