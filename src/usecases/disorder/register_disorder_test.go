package disorder_test

import (
	"pesthub/adapters/memdb"
	"pesthub/adapters/testmsgs"
	"pesthub/env"
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

// benchmark
var names100k []string

func init() {
	// prepare names for benchmarking
	names100k = make([]string, 100_000)
	for i := 0; i < 100_000; i++ {
		names100k[i] = "randomname" + strconv.Itoa(i)
	}

	// prepare env
	configureEnv()
}

func BenchmarkRegisterDisorder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := disorder.RegisterDisorderInput{
			Name: names100k[i],
		}
		disorder.RegisterDisorder(&input)
	}
}
