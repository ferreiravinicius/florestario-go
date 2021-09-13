package disorder_test

import (
	"pesthub/contracts/store"
	"pesthub/entities"
	"pesthub/usecases/disorder"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockFindDisordersHavingNames(resOk []*entities.Disorder, resErr error) store.FindDisorderHavingNames {
	return func(names ...string) ([]*entities.Disorder, error) {
		return resOk, resErr
	}
}

func TestCheckAlreadyExists(t *testing.T) {

	t.Run("it should return error when name already exists", func(t *testing.T) {
		wantedResult := make([]*entities.Disorder, 1) // 1 result
		mockFindAll := mockFindDisordersHavingNames(wantedResult, nil)
		fnCheckExists := disorder.NewCheckAlreadyExists(mockFindAll)
		input := &disorder.CheckAlreadyExistsInput{}
		err := fnCheckExists(input)
		assert.Error(t, err)
	})

	t.Run("it should pass when no duplicated is found", func(t *testing.T) {
		wantedResult := make([]*entities.Disorder, 0) // empty result
		mockFindAll := mockFindDisordersHavingNames(wantedResult, nil)
		fnCheckExists := disorder.NewCheckAlreadyExists(mockFindAll)
		input := &disorder.CheckAlreadyExistsInput{}
		err := fnCheckExists(input)
		assert.NoError(t, err)
	})

	t.Run("it should allow exclude", func(t *testing.T) {
		// TODO: there's no id defined yet
	})

}
