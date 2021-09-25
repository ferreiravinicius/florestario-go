package testmsgs_test

import (
	"pesthub/adapters/testmsgs"
	"pesthub/contracts"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetText(t *testing.T) {
	provider := testmsgs.NewTestableMessages()
	msg := provider.GetText("message.mock")
	assert.Equal(t, "message.mock", msg)

	_ = provider.GetText("text", contracts.ArgMap{"arg": "value"})

}
