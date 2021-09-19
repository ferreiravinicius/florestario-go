package failure_test

import (
	"pesthub/failure"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseCaseError(t *testing.T) {

	e1 := failure.UseCaseError{Message: "message"}
	assert.Error(t, e1)
	assert.Equal(t, "message", e1.Message)

	e2 := failure.UseCase("message")
	assert.Error(t, e2)
	assert.Equal(t, "message", e2.Message)
	assert.NotEmpty(t, e2.Error())
}
