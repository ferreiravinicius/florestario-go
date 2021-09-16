package disorder_test

import (
	"pesthub/commons/errors"
	"pesthub/contracts/store"
	"pesthub/entities"
	"pesthub/usecases/disorder"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockExistsByName(resOk bool, resErr error) store.ExistsDisorderByName {
	return func(_ string) (bool, error) {
		return resOk, resErr
	}
}

func mockSave(rsId int64, rsErr error) store.SaveDisorder {
	return func(_ *entities.Disorder) (int64, error) {
		return rsId, rsErr
	}
}

func TestCreate(t *testing.T) {

	var getValidInput = func() *disorder.CreateDisorderInput {
		return &disorder.CreateDisorderInput{
			Name: "name",
		}
	}

	t.Run("it should return generated unique field", func(t *testing.T) {
		create := disorder.NewCreateDisorder(mockSave(999, nil), mockExistsByName(false, nil))
		id, err := create(getValidInput())
		assert.NoError(t, err)
		assert.Equal(t, int64(999), id)
	})

	t.Run("it should return error when invalid input is provided", func(t *testing.T) {
		input := disorder.CreateDisorderInput{
			Name: "", // invalid
		}
		create := disorder.NewCreateDisorder(mockSave(999, nil), mockExistsByName(false, nil))
		_, err := create(&input)
		assert.Error(t, err)
	})

	t.Run("it should return error when found duplicated", func(t *testing.T) {
		create := disorder.NewCreateDisorder(mockSave(999, nil), mockExistsByName(true, nil))
		_, err := create(getValidInput())
		assert.Error(t, err)
		assert.IsType(t, errors.BusinessError{}, err)
	})

}
