package pests_test

import (
	"errors"
	"pesthub/entities"
	"pesthub/features/pests"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockInsertWithError pests.InsertCommand = func(_ *entities.Pest) (int64, error) {
	return 0, errors.New("fake error")
}

var mockInsert pests.InsertCommand = func(_ *entities.Pest) (int64, error) {
	return 999, nil
}

func TestCreate(t *testing.T) {

	var getValidInput = func() pests.CreateInput {
		return pests.CreateInput{CommonName: "testing"}
	}

	t.Run("it should return generated unique field", func(t *testing.T) {
		input := getValidInput()
		deps := pests.CreateDeps{InsertCommand: mockInsert}
		id, err := pests.Create(deps, &input)
		assert.NoError(t, err)
		assert.Equal(t, int64(999), id)
	})

	t.Run("it should return error when invalid input is provided", func(t *testing.T) {
		wrongInput := pests.CreateInput{CommonName: ""}
		deps := pests.CreateDeps{InsertCommand: mockInsert}
		_, err := pests.Create(deps, &wrongInput)
		assert.Error(t, err)
	})

	t.Run("it should return error when insert fails for whatever reason", func(t *testing.T) {
		input := getValidInput()
		deps := pests.CreateDeps{InsertCommand: mockInsertWithError}
		_, err := pests.Create(deps, &input)
		assert.Error(t, err)
	})
}
