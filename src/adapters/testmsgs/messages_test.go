package testmsgs_test

import (
	"pesthub/adapters/testmsgs"
	"pesthub/contracts"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetText(t *testing.T) {
	provider := testmsgs.NewTestableMessageProvider()
	msg := provider.Get("message.mock")
	assert.Equal(t, "message.mock", msg)

	_ = provider.Get("text", contracts.ArgMap{"arg": "value"})

}
