package failures_test

import (
	"pesthub/failures"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationError(t *testing.T) {

	v := failures.ValidationError{
		Field:   "field",
		Message: "message",
	}
	assert.Error(t, v)

	err1 := failures.Validation("message")
	assert.IsType(t, failures.ValidationError{}, err1)
	assert.Equal(t, "message", err1.Message)
	assert.Equal(t, "", err1.Field)

	err2 := failures.Validation("message", "field")
	assert.Equal(t, "message", err2.Message)
	assert.Equal(t, "field", err2.Field)

}
