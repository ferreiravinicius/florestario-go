package errors_test

import (
	"pesthub/commons/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBusinessError(t *testing.T) {
	got := errors.Business("code")
	assert.Error(t, got)
	assert.IsType(t, errors.BusinessError{}, got)
	assert.Equal(t, "code", got.Code)
	assert.Nil(t, got.Args)

	gotArgs := errors.Business("code", map[string]string{"a": "b"})
	args := gotArgs.Args
	assert.NotNil(t, args)
	assert.Contains(t, args, "a")
}
