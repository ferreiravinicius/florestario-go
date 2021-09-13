package disorder_test

import (
	"errors"
	"pesthub/contracts/store"
	"pesthub/entities"
	"pesthub/usecases/disorder"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockCheckExists(result error) disorder.ICheckAlreadyExists {
	return func(data *disorder.CheckAlreadyExistsInput) error {
		return result
	}
}

func mockSave(rsId int64, rsErr error) store.SaveDisorder {
	return func(_ *entities.Disorder) (int64, error) {
		return rsId, rsErr
	}
}

func TestCreate(t *testing.T) {

	var getValidInput = func() disorder.CreateDisorderInput {
		return disorder.CreateDisorderInput{
			Name:          "name",
			BionomialName: "binomial name",
		}
	}

	t.Run("it should return generated unique field", func(t *testing.T) {
		input := getValidInput()
		create := disorder.NewCreateDisorder(mockSave(999, nil), mockCheckExists(nil))
		id, err := create(&input)
		assert.NoError(t, err)
		assert.Equal(t, int64(999), id)
	})

	t.Run("it should return error when invalid input is provided", func(t *testing.T) {
		input := disorder.CreateDisorderInput{
			Name:          "", // invalid
			BionomialName: "", // invalid
		}
		create := disorder.NewCreateDisorder(mockSave(999, nil), mockCheckExists(nil))
		_, err := create(&input)
		assert.Error(t, err)
	})

	t.Run("it should return error when found duplicated", func(t *testing.T) {
		input := getValidInput()
		wanted := errors.New("fake error")
		create := disorder.NewCreateDisorder(mockSave(999, nil), mockCheckExists(wanted))
		_, err := create(&input)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), wanted.Error())
	})

}
