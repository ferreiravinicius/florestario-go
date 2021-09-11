package pest_test

import (
	"errors"
	"pesthub/contracts/store"
	"pesthub/entities"
	"testing"
)

// import (
// 	"errors"
// 	"pesthub/contracts"
// 	"pesthub/entities"
// 	"pesthub/usecases/pest"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

func mockSave(rsId int64, rsErr error) store.SavePest {
	return func(_ *entities.Pest) (int64, error) {
		return rsId, rsErr
	}
}

var (
	mockSaveSuccess = mockSave(999, nil)
	mockSaveErr     = mockSave(0, errors.New("random error"))
)

func TestCreate(t *testing.T) {

	// 	var getValidInput = func() pest.CreatePestInput {
	// 		return pest.CreatePestInput{Name: "testing"}
	// 	}

	// 	t.Run("it should return generated unique field", func(t *testing.T) {
	// 		input := getValidInput()
	// 		usecase := pest.NewCreatePest(mockSaveSuccess, mockGetByNameNoResult)
	// 		id, err := usecase.Execute(&input)
	// 		assert.NoError(t, err)
	// 		assert.Equal(t, int64(999), id)
	// 	})

	// 	t.Run("it should return error when invalid input is provided", func(t *testing.T) {
	// 		wrongInput := pest.CreatePestInput{Name: ""}
	// 		usecase := pest.NewCreatePest(mockSaveSuccess, mockGetByNameNoResult)
	// 		_, err := usecase.Execute(&wrongInput)
	// 		assert.Error(t, err)
	// 	})

	// 	t.Run("it should return error when insert fails for whatever reason", func(t *testing.T) {
	// 		input := getValidInput()
	// 		usecase := pest.NewCreatePest(mockSaveErr, mockGetByNameNoResult)
	// 		_, err := usecase.Execute(&input)
	// 		assert.Error(t, err)
	// 	})

	// 	t.Run("it should return error when name already exists", func(t *testing.T) {
	// 		input := getValidInput()
	// 		usecase := pest.NewCreatePest(mockSaveErr, mockGetByNameHavingResult)
	// 		_, err := usecase.Execute(&input)
	// 		assert.Error(t, err)
	// 		assert.ErrorIs(t, err, pest.ErrDuplicated) // may change
	// 	})

}
