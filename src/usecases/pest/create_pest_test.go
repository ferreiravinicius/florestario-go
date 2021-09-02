package pest_test

import (
	"errors"
	"pesthub/contracts"
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

func mockGetByName(rsPest *entities.Pest, rsErr error) contracts.GetPestByName {
	return func(_ string) (*entities.Pest, error) {
		return rsPest, rsErr
	}
}

var (
	mockGetByNameNoResult     = mockGetByName(nil, nil)
	mockGetByNameHavingResult = mockGetByName(&entities.Pest{}, nil)
)

func TestCreate(t *testing.T) {

	var getValidInput = func() pest.CreatePestInput {
		return pest.CreatePestInput{CommonName: "testing"}
	}

	t.Run("it should return generated unique field", func(t *testing.T) {
		input := getValidInput()
		usecase := pest.NewCreatePest(mockInsert, mockGetByNameNoResult)
		id, err := usecase.Execute(&input)
		assert.NoError(t, err)
		assert.Equal(t, int64(999), id)
	})

	t.Run("it should return error when invalid input is provided", func(t *testing.T) {
		wrongInput := pest.CreatePestInput{CommonName: ""}
		usecase := pest.NewCreatePest(mockInsert, mockGetByNameNoResult)
		_, err := usecase.Execute(&wrongInput)
		assert.Error(t, err)
	})

	t.Run("it should return error when insert fails for whatever reason", func(t *testing.T) {
		input := getValidInput()
		usecase := pest.NewCreatePest(mockInsertWithError, mockGetByNameNoResult)
		_, err := usecase.Execute(&input)
		assert.Error(t, err)
	})
	// TODO: rewrite tests for duplicated name validation
}
