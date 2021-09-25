package validators_test

import (
	"pesthub/adapters/testmsgs"
	"pesthub/env"
	"pesthub/validators"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	mp := testmsgs.NewTestableMessageProvider()
	env.MessageProvider = mp
}

func TestName(t *testing.T) {

	t.Run("name must have atleast 3 chars ", func(t *testing.T) {
		err := validators.Name("ab")
		assert.NotNil(t, err)
		wanted := env.MessageProvider.Get(validators.MsgNameMinimumSize)
		assert.Equal(t, wanted, err.Error())
	})

	t.Run("should pass when provided valid name", func(t *testing.T) {
		err := validators.Name("valid name")
		assert.Nil(t, err)
	})
}
