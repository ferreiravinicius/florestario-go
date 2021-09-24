package disorder_test

import (
	"pesthub/adapters/memdb"
	"pesthub/adapters/testmsgs"
	"pesthub/failures"
	"pesthub/usecases/disorder"
	"strconv"
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
	t.Run("it should return generated id", func(t *testing.T) {
		store := memdb.NewMemoryDisorderStore()
		messages := testmsgs.NewTestableMessages()
		sut := disorder.NewRegisterDisorder(store, messages)

		input := makeTestableRegisterDisorderInput()
		output, err := sut.Execute(input)
		assert.NoError(t, err)
		assert.NotEmpty(t, output.Id)
	})

	t.Run("it should return error when name already exists", func(t *testing.T) {
		store := memdb.NewMemoryDisorderStore()
		messages := testmsgs.NewTestableMessages()
		sut := disorder.NewRegisterDisorder(store, messages)

		input := makeTestableRegisterDisorderInput()
		entity := input.ToEntity()
		store.Save(entity)

		_, err := sut.Execute(input)
		assert.Error(t, err)
		assert.IsType(t, failures.UseCaseError{}, err)
		assert.Equal(t, disorder.MsgNameAlreadyExists, err.Error())
	})
}

// benchmark
var names100k []string

func init() {
	// prepare names for benchmarking
	names100k = make([]string, 100_000)
	for i := 0; i < 100_000; i++ {
		names100k[i] = "randomname" + strconv.Itoa(i)
	}
}

func BenchmarkRegisterDisorder(b *testing.B) {
	sut := disorder.NewRegisterDisorder(memdb.NewMemoryDisorderStore(), testmsgs.NewTestableMessages())
	for i := 0; i < b.N; i++ {
		input := disorder.RegisterDisorderInput{
			Name: names100k[i],
		}
		sut.Execute(&input)
	}
}
