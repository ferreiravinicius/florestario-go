package pest_test

import (
	"errors"
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

	t.Run("it should return error when found duplicated pest", func(t *testing.T) {
		input := getValidInput()
		wanted := errors.New("fake error")
		create := pest.NewCreatePest(mockSave(999, nil), mockCheckExists(wanted))
		_, err := create(&input)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), wanted.Error())
	})

}
