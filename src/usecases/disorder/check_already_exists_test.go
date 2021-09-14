package disorder_test

import (
	"pesthub/contracts/store"
	"pesthub/entities"
	"pesthub/usecases/disorder"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockGetDisorderByName(resOk *entities.Disorder, resErr error) store.GetDisorderByName {
	return func(_ string) (*entities.Disorder, error) {
		return resOk, resErr
	}
}

func TestCheckAlreadyExists(t *testing.T) {

	t.Run("it should return error when name already exists", func(t *testing.T) {
		wantedResult := &entities.Disorder{}
		mockGetDisorder := mockGetDisorderByName(wantedResult, nil)
		check := disorder.NewCheckAlreadyExists(mockGetDisorder)
		input := &disorder.CheckAlreadyExistsInput{}
		err := check(input)
		assert.Error(t, err)
	})

	t.Run("it should pass when no duplicated name is found", func(t *testing.T) {
		mockGetDisorder := mockGetDisorderByName(nil, nil)
		check := disorder.NewCheckAlreadyExists(mockGetDisorder)
		input := &disorder.CheckAlreadyExistsInput{}
		err := check(input)
		assert.NoError(t, err)
	})

	t.Run("it should allow exclude", func(t *testing.T) {
		// TODO: there's no id defined yet
	})

}
