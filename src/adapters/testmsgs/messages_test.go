package testmsgs_test

import (
	"pesthub/adapters/testmsgs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetText(t *testing.T) {
	provider := testmsgs.NewTestableMessages()
	msg := provider.GetText("message.mock")
	assert.Equal(t, "message.mock", msg)
}
