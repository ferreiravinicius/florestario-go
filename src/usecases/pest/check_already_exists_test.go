package pest_test

import (
	"pesthub/contracts/store"
	"pesthub/entities"
	"pesthub/usecases/pest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockFindPestsHavingNames(resOk []*entities.Pest, resErr error) store.FindPestsHavingNames {
	return func(names ...string) ([]*entities.Pest, error) {
		return resOk, resErr
	}
}

func TestCheckAlreadyExists(t *testing.T) {

	t.Run("it should return error when name already exists", func(t *testing.T) {
		wantedResult := make([]*entities.Pest, 1) // 1 result
		mockFindAll := mockFindPestsHavingNames(wantedResult, nil)
		fnCheckExists := pest.NewCheckAlreadyExists(mockFindAll)
		input := &pest.CheckAlreadyExistsInput{}
		err := fnCheckExists(input)
		assert.Error(t, err)
	})

	t.Run("it should pass when no duplicated is found", func(t *testing.T) {
		wantedResult := make([]*entities.Pest, 0) // empty result
		mockFindAll := mockFindPestsHavingNames(wantedResult, nil)
		fnCheckExists := pest.NewCheckAlreadyExists(mockFindAll)
		input := &pest.CheckAlreadyExistsInput{}
		err := fnCheckExists(input)
		assert.NoError(t, err)
	})

	t.Run("it should allow exclude", func(t *testing.T) {
		// there's no id defined yet
	})

}
