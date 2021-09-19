package disorder_test

import (
	"pesthub/adapters/memdb"
	"pesthub/env"
	"testing"
)

func init() {
	env.DisorderStore = memdb.NewMemoryDisorderStore()
	// env.Messages = simplemsgs.
}

func TestRegisterDisorder(t *testing.T) {
	t.Run("", func(t *testing.T) {

	})
}
