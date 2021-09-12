package pest_test

import (
	"pesthub/contracts/store"
	"pesthub/entities"
	"pesthub/usecases/pest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockCheckExists(result error) pest.CheckAlreadyExists {
	return func(data *pest.CheckAlreadyExistsInput) error {
		return result
	}
}

func mockSave(rsId int64, rsErr error) store.SavePest {
	return func(_ *entities.Pest) (int64, error) {
		return rsId, rsErr
	}
}

func TestCreate(t *testing.T) {

	var getValidInput = func() pest.CreatePestInput {
		return pest.CreatePestInput{
			Name:          "name",
			BionomialName: "binomial name",
		}
	}

	t.Run("it should return generated unique field", func(t *testing.T) {
		input := getValidInput()
		create := pest.NewCreatePest(mockSave(999, nil), mockCheckExists(nil))
		id, err := create(&input)
		assert.NoError(t, err)
		assert.Equal(t, int64(999), id)
	})

	t.Run("it should return error when invalid input is provided", func(t *testing.T) {
		input := pest.CreatePestInput{
			Name:          "", // invalid
			BionomialName: "", // invalid
		}
		create := pest.NewCreatePest(mockSave(999, nil), mockCheckExists(nil))
		_, err := create(&input)
		assert.Error(t, err)
	})
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
