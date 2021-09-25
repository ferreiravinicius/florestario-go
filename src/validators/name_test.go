package validators_test

import (
	"pesthub/adapters/testmsgs"
	"pesthub/validators"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	messages := testmsgs.NewTestableMessages()

	t.Run("name must have atleast 3 chars ", func(t *testing.T) {
		err1 := validators.Name(messages, "ab")
		assert.NotNil(t, err1)

		wanted := messages.GetText(validators.MsgNameMinimumSize)
		assert.Equal(t, wanted, err1.Error())

		err2 := validators.Name(messages, "valid name")
		assert.Nil(t, err2)
	})

}
