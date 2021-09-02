package pest_test

import (
	"errors"
	"pesthub/entities"
	"pesthub/usecases/pest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockInsertWithError = func(_ *entities.Pest) (int64, error) {
	return 0, errors.New("fake error")
}

var mockInsert = func(_ *entities.Pest) (int64, error) {
	return 999, nil
}

func TestCreate(t *testing.T) {

	var getValidInput = func() pest.CreatePestInput {
		return pest.CreatePestInput{CommonName: "testing"}
	}

	t.Run("it should return generated unique field", func(t *testing.T) {
		input := getValidInput()
		usecase := pest.NewCreatePest(mockInsert)
		id, err := usecase.Execute(&input)
		assert.NoError(t, err)
		assert.Equal(t, int64(999), id)
	})

	t.Run("it should return error when invalid input is provided", func(t *testing.T) {
		wrongInput := pest.CreatePestInput{CommonName: ""}
		usecase := pest.NewCreatePest(mockInsert)
		_, err := usecase.Execute(&wrongInput)
		assert.Error(t, err)
	})

	t.Run("it should return error when insert fails for whatever reason", func(t *testing.T) {
		input := getValidInput()
		usecase := pest.NewCreatePest(mockInsertWithError)
		_, err := usecase.Execute(&input)
		assert.Error(t, err)
	})
}
