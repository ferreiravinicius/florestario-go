package pest_test

import (
	"errors"
	"pesthub/entities"
	"pesthub/features/pest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockInsertWithError pest.InsertCommand = func(_ *entities.Pest) (int64, error) {
	return 0, errors.New("fake error")
}

var mockInsert pest.InsertCommand = func(_ *entities.Pest) (int64, error) {
	return 999, nil
}

func TestCreate(t *testing.T) {

	var getValidInput = func() pest.CreateInput {
		return pest.CreateInput{CommonName: "testing"}
	}

	t.Run("it should return generated unique field", func(t *testing.T) {
		input := getValidInput()
		deps := pest.CreateDeps{InsertCommand: mockInsert}
		id, err := pest.Create(deps, &input)
		assert.NoError(t, err)
		assert.Equal(t, int64(999), id)
	})

	t.Run("it should return error when invalid input is provided", func(t *testing.T) {
		wrongInput := pest.CreateInput{CommonName: ""}
		deps := pest.CreateDeps{InsertCommand: mockInsert}
		_, err := pest.Create(deps, &wrongInput)
		assert.Error(t, err)
	})

	t.Run("it should return error when insert fails for whatever reason", func(t *testing.T) {
		input := getValidInput()
		deps := pest.CreateDeps{InsertCommand: mockInsertWithError}
		_, err := pest.Create(deps, &input)
		assert.Error(t, err)
	})
}
