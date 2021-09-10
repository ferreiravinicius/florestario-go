package pest_test

import (
	"pesthub/entities"
	"pesthub/usecases/pest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockFindAllHavingNames(resOk []*entities.Pest, resErr error) pest.FindAllHavingNames {
	return func(names ...string) ([]*entities.Pest, error) {
		return resOk, resErr
	}
}

func TestCheckAlreadyExists(t *testing.T) {

	t.Run("should return error when name already exists", func(t *testing.T) {
		mockResult := []*entities.Pest{&entities.Pest{}} // 1 result
		mockFindAll := mockFindAllHavingNames(mockResult, nil)
		fnCheckExists := pest.NewCheckAlreadyExists(mockFindAll)
		input := &pest.CheckAlreadyExistsInput{}
		err := fnCheckExists(input)
		assert.Error(t, err)
	})

	t.Run("should pass when theres no duplicated values", func(t *testing.T) {
		mockResult := make([]*entities.Pest, 0) // empty result
		mockFindAll := mockFindAllHavingNames(mockResult, nil)
		fnCheckExists := pest.NewCheckAlreadyExists(mockFindAll)
		input := &pest.CheckAlreadyExistsInput{}
		err := fnCheckExists(input)
		assert.NoError(t, err)
	})
	//should validate common name
	//should validate scientific name

}
