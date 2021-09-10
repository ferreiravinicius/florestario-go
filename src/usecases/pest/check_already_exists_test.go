package pest_test

import (
	"pesthub/contracts"
	"pesthub/entities"
	"pesthub/usecases/pest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockFindPestsHavingNames(resOk []*entities.Pest, resErr error) contracts.FindPestsHavingNames {
	return func(names ...string) ([]*entities.Pest, error) {
		return resOk, resErr
	}
}

func TestCheckAlreadyExists(t *testing.T) {

	t.Run("it should return error when name already exists", func(t *testing.T) {
		desiredResult := []*entities.Pest{&entities.Pest{}} // 1 result
		mockFindAll := mockFindPestsHavingNames(desiredResult, nil)
		fnCheckExists := pest.NewCheckAlreadyExists(mockFindAll)
		input := &pest.CheckAlreadyExistsInput{}
		err := fnCheckExists(input)
		assert.Error(t, err)
	})

	t.Run("it should pass when no duplicated is found", func(t *testing.T) {
		desiredResult := make([]*entities.Pest, 0) // empty result
		mockFindAll := mockFindPestsHavingNames(desiredResult, nil)
		fnCheckExists := pest.NewCheckAlreadyExists(mockFindAll)
		input := &pest.CheckAlreadyExistsInput{}
		err := fnCheckExists(input)
		assert.NoError(t, err)
	})

}
